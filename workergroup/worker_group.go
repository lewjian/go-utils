package workergroup

import (
	"runtime"
	"sync"

	"github.com/lewjian/utils/rescue"
)

type WorkerGroup struct {
	c  chan struct{}
	wg sync.WaitGroup
}

// Run a task
func (tg *WorkerGroup) Run(f func()) {
	// 检查是否可以执行
	tg.c <- struct{}{}
	tg.wg.Add(1)
	go func() {
		defer rescue.Recover(func() {
			tg.wg.Done()
			<-tg.c
		})
		f()
	}()
}

// Wait all task finish
func (tg *WorkerGroup) Wait() {
	tg.wg.Wait()
}

// New a WorkerGroup object
func New(workerNum int) *WorkerGroup {
	if workerNum <= 0 {
		workerNum = runtime.NumCPU()
	}
	return &WorkerGroup{
		c: make(chan struct{}, workerNum),
	}
}
