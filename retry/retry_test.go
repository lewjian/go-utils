package retry

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	r := NewRetry()
	count := 0
	_ = r.Run(func() error {
		count++
		fmt.Printf("excute times: %d\n", count)
		return errors.New("continue execute")
	})
	assert.NotNil(t, defaultRetryTimes, count)
}

func TestInfinite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	r := NewRetry(WithInfiniteRetry())
	count := 0
	err := r.Run(func() error {
		count++
		fmt.Printf("excute times: %d\n", count)
		return errors.New("continue execute")
	})
	assert.NotNil(t, err)
}

func TestWithRetryNum(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	retryNum := 10
	r := NewRetry(WithRetryNum(retryNum))
	count := 0
	_ = r.Run(func() error {
		count++
		fmt.Printf("excute times: %d\n", count)
		return errors.New("continue execute")
	})
	assert.Equal(t, retryNum, count)
}

func TestWithIncrFactor(t *testing.T) {
	r := NewRetry(WithIncrFactor(2))
	now := time.Now()
	end := time.Now()
	count := 0
	r.Run(func() error {
		end = time.Now()
		count++
		fmt.Printf("%s: excute times: %d\n", end.String(), count)
		return errors.New("continue execute")
	})
	assert.Equal(t, 4, int(end.Sub(now).Seconds()))
}

func TestWithContext(t *testing.T) {
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	ctx, cancel := context.WithCancel(context.Background())
	r := NewRetry(WithIncrFactor(2), WithContext(ctx))
	count := 0
	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()
	err := r.Run(func() error {
		count++
		fmt.Printf("%s: excute times: %d\n", time.Now().String(), count)
		return errors.New("continue execute")
	})
	assert.Equal(t, ErrContextCancelled, err)
}

func TestWithInterval(t *testing.T) {
	r := NewRetry(WithInterval(time.Second * 2))
	now := time.Now()
	end := time.Now()
	count := 0
	r.Run(func() error {
		end = time.Now()
		count++
		fmt.Printf("%s: excute times: %d\n", end.String(), count)
		return errors.New("continue execute")
	})
	assert.Equal(t, 4, int(end.Sub(now).Seconds()))
}

func TestComplex(t *testing.T) {
	r := NewRetry(WithInterval(time.Second*2), WithIncrFactor(1), WithRetryNum(5))
	now := time.Now()
	end := time.Now()
	count := 0 // 2 4 8 16
	r.Run(func() error {
		end = time.Now()
		count++
		fmt.Printf("%s: excute times: %d\n", end.String(), count)
		return errors.New("continue execute")
	})
	assert.Equal(t, 30, int(end.Sub(now).Seconds()))
}
