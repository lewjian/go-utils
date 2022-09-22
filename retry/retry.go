package retry

import (
	"context"
	"errors"
	"time"
)

const (
	defaultRetryTimes = 3
	defaultInterval   = time.Second
)

var ErrContextCancelled = errors.New("context cancelled or timeout")

type Retry struct {
	incrFactor   float64       // 执行时间间隙增长因子
	interval     time.Duration // 基础时间间隙
	retryNum     int           // 重试次数
	executeCount int           // 累计执行次数
	ctx          context.Context
}
type Option func(retry *Retry)

// WithInterval 设置间隙等待时间
func WithInterval(interval time.Duration) Option {
	return func(r *Retry) {
		r.interval = interval
	}
}

// WithIncrFactor 设置间隙增长因子
func WithIncrFactor(factor float64) Option {
	return func(r *Retry) {
		r.incrFactor = factor
	}
}

// WithRetryNum 设置执行次数
func WithRetryNum(retryNum int) Option {
	return func(r *Retry) {
		r.retryNum = retryNum
	}
}

// WithInfiniteRetry 无限执行
func WithInfiniteRetry() Option {
	return func(r *Retry) {
		r.retryNum = -1
	}
}

// WithContext 设置context
func WithContext(ctx context.Context) Option {
	return func(r *Retry) {
		r.ctx = ctx
	}
}

// Run a retry task
func (r *Retry) Run(f func() error) error {
	var err error
	for r.retryNum <= 0 || (r.executeCount < r.retryNum) {
		select {
		case <-r.ctx.Done():
			return ErrContextCancelled
		default:
		}
		r.executeCount++
		err = f()
		if err == nil {
			return nil
		}
		if r.executeCount == r.retryNum {
			return err
		}
		// 休眠一下
		time.Sleep(r.interval)

		// 重新计算间隙interval
		r.interval = time.Duration(int64(float64(r.interval) * (1 + r.incrFactor)))
	}
	return err
}

// NewRetry create a new Retry
func NewRetry(opts ...Option) *Retry {
	r := &Retry{
		interval: defaultInterval,
		retryNum: defaultRetryTimes,
		ctx:      context.Background(),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
