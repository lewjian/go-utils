package set

import (
	"sort"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetInt(t *testing.T) {
	s := NewSet[int]()
	var wg sync.WaitGroup
	want := make([]int, 100)
	for i := 0; i < 100; i++ {
		want[i] = i
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(i int) {
				s.Add(i)
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
	got := s.Values()
	sort.Ints(got)
	assert.Equal(t, want, got)
}

func TestSetString(t *testing.T) {
	s := NewSet[string]()
	var wg sync.WaitGroup
	want := make([]string, 100)
	for i := 0; i < 100; i++ {
		str := strconv.Itoa(i)
		want[i] = str
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(str string) {
				s.Add(str)
				wg.Done()
			}(str)
		}
	}
	wg.Wait()
	got := s.Values()
	sort.Strings(got)
	sort.Strings(want)
	assert.Equal(t, want, got)
}

func TestSetDel(t *testing.T) {
	s := NewSet[string]()
	var wg sync.WaitGroup
	want := make([]string, 0, 100)
	for i := 0; i < 100; i++ {
		str := strconv.Itoa(i)
		if i >= 50 {
			want = append(want, str)
		}
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(str string) {
				s.Add(str)
				wg.Done()
			}(str)
		}
	}
	for i := 0; i < 50; i++ {
		s.Del(strconv.Itoa(i))
	}
	wg.Wait()
	got := s.Values()
	sort.Strings(got)
	sort.Strings(want)
	assert.Equal(t, want, got)
}
