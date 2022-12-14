package bloom

// boom copy from github.com/zeromicro/go-zero/core/bloom
import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spaolacci/murmur3"
)

const (
	// for detailed error rate table, see http://pages.cs.wisc.edu/~cao/papers/summary-cache/node8.html
	// maps as k in the error rate table
	maps      = 14
	setScript = `
for _, offset in ipairs(ARGV) do
	redis.call("setbit", KEYS[1], offset, 1)
end
`
	testScript = `
for _, offset in ipairs(ARGV) do
	if tonumber(redis.call("getbit", KEYS[1], offset)) == 0 then
		return false
	end
end
return true
`
)

// ErrTooLargeOffset indicates the offset is too large in bitset.
var ErrTooLargeOffset = errors.New("too large offset")

type (
	// A Filter is a bloom filter.
	Filter struct {
		bits   uint
		bitSet bitSetProvider
	}

	bitSetProvider interface {
		check(context.Context, []uint) (bool, error)
		set(context.Context, []uint) error
	}
)

// New create a Filter, store is the backed redis, key is the key for the bloom filter,
// bits is how many bits will be used, maps is how many hashes for each addition.
// best practices:
// elements - means how many actual elements
// when maps = 14, formula: 0.7*(bits/maps), bits = 20*elements, the error rate is 0.000067 < 1e-4
// for detailed error rate table, see http://pages.cs.wisc.edu/~cao/papers/summary-cache/node8.html
func New(store *redis.Client, key string, bits uint) *Filter {
	return &Filter{
		bits:   bits,
		bitSet: newRedisBitSet(store, key, bits),
	}
}

// Add adds data into f.
func (f *Filter) Add(ctx context.Context, data []byte) error {
	locations := f.getLocations(data)
	return f.bitSet.set(ctx, locations)
}

// Exists checks if data is in f.
func (f *Filter) Exists(ctx context.Context, data []byte) (bool, error) {
	locations := f.getLocations(data)
	isSet, err := f.bitSet.check(ctx, locations)
	if err != nil {
		return false, err
	}

	return isSet, nil
}

func (f *Filter) getLocations(data []byte) []uint {
	locations := make([]uint, maps)
	for i := uint(0); i < maps; i++ {
		hashValue := murmur3.Sum64(append(data, byte(i)))
		locations[i] = uint(hashValue % uint64(f.bits))
	}

	return locations
}

type redisBitSet struct {
	store *redis.Client
	key   string
	bits  uint
}

func newRedisBitSet(store *redis.Client, key string, bits uint) *redisBitSet {
	return &redisBitSet{
		store: store,
		key:   key,
		bits:  bits,
	}
}

func (r *redisBitSet) buildOffsetArgs(offsets []uint) ([]string, error) {
	var args []string

	for _, offset := range offsets {
		if offset >= r.bits {
			return nil, ErrTooLargeOffset
		}

		args = append(args, strconv.FormatUint(uint64(offset), 10))
	}

	return args, nil
}

func (r *redisBitSet) check(ctx context.Context, offsets []uint) (bool, error) {
	args, err := r.buildOffsetArgs(offsets)
	if err != nil {
		return false, err
	}
	cmd := r.store.Eval(ctx, testScript, []string{r.key}, args)
	if cmd.Err() == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	resp, err := cmd.Result()
	exists, ok := resp.(int64)
	if !ok {
		return false, nil
	}

	return exists == 1, nil
}

func (r *redisBitSet) del(ctx context.Context) error {
	res := r.store.Del(ctx, r.key)
	return res.Err()
}

func (r *redisBitSet) expire(ctx context.Context, seconds int) error {
	res := r.store.Expire(ctx, r.key, time.Second*time.Duration(seconds))
	return res.Err()
}

func (r *redisBitSet) set(ctx context.Context, offsets []uint) error {
	args, err := r.buildOffsetArgs(offsets)
	if err != nil {
		return err
	}

	cmd := r.store.Eval(ctx, setScript, []string{r.key}, args)
	if cmd.Err() == redis.Nil {
		return nil
	}

	return err
}
