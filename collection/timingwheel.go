package collection

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lewjian/go-utils/number"
	"github.com/lewjian/go-utils/rescue"
)

type (
	TimingWheel struct {
		interval     time.Duration
		numSlot      int
		currentSlot  int
		lock         sync.RWMutex
		ticker       *time.Ticker
		lastTickTime time.Time
		started      int32
		stopChan     chan struct{}
		tasks        []*Task
	}
	Task struct {
		ID          string
		ExecuteTime time.Time
		Execute     Execute
		Next        *Task
	}
	Option  func(tw *TimingWheel)
	Execute func(ID string)
)

var (
	ErrInvalidDuration = errors.New("duration must greater than interval")
	ErrInvalidTask     = errors.New("invalid task")
)

func (tw *TimingWheel) tick() {
	if !atomic.CompareAndSwapInt32(&tw.started, 0, 1) {
		return
	}
	for {
		_, ok := <-tw.ticker.C
		if !ok {
			break
		}

		tw.currentSlot = (tw.currentSlot + 1) % tw.numSlot
		tw.lastTickTime = time.Now()
		go tw.scanAndRunTasks()
	}
}

// AddTask adds tasks
func (tw *TimingWheel) AddTask(ID string, duration time.Duration, execute Execute) error {
	if execute == nil {
		return ErrInvalidTask
	}
	if duration < tw.interval {
		return ErrInvalidDuration
	}
	task := tw.newTask(ID, duration, execute)
	tw.lock.Lock()
	defer tw.lock.Unlock()
	slot := (number.RoundInt(float64(duration)/float64(tw.interval)) + tw.currentSlot) % tw.numSlot
	oldTask := tw.tasks[slot]
	tw.tasks[slot] = insertTask(task, oldTask)
	return nil
}

func insertTask(task, root *Task) *Task {
	if task == nil {
		return root
	}
	if root == nil {
		return task
	}
	head := &Task{
		Next: root,
	}
	prev := head
	for root != nil {
		if root.ExecuteTime.After(task.ExecuteTime) {
			break
		}
		prev = prev.Next
		root = root.Next
	}
	prev.Next = task
	task.Next = root
	return head.Next
}

// Stop timing wheel
func (tw *TimingWheel) Stop() {
	tw.ticker.Stop()
	close(tw.stopChan)
}

// Wait until tw stop
func (tw *TimingWheel) Wait() bool {
	<-tw.stopChan
	return true
}

func (tw *TimingWheel) scanAndRunTasks() {
	slot := tw.currentSlot
	task := tw.tasks[slot]
	tickTime := tw.lastTickTime
	for task != nil {
		if !task.ExecuteTime.Before(tickTime.Add(tw.interval)) {
			break
		}
		t := task
		go rescue.RunSafe(func() {
			t.Execute(t.ID)
		})
		task = task.Next
	}
	tw.lock.Lock()
	tw.tasks[slot] = task
	tw.lock.Unlock()
}

func (tw *TimingWheel) newTask(ID string, duration time.Duration, execute Execute) *Task {
	if execute == nil {
		return nil
	}
	return &Task{
		ID:          ID,
		Next:        nil,
		ExecuteTime: tw.lastTickTime.Add(duration),
		Execute:     execute,
	}
}

// NewTimingWheel creates new TimingWheel
func NewTimingWheel(slots int, tickInterval time.Duration) *TimingWheel {
	if slots <= 1 {
		panic("invalid parameter")
	}
	tw := &TimingWheel{
		numSlot:      slots,
		currentSlot:  slots - 1,
		interval:     tickInterval,
		stopChan:     make(chan struct{}),
		ticker:       time.NewTicker(tickInterval),
		lastTickTime: time.Now(),
		tasks:        make([]*Task, slots),
	}
	go tw.tick()
	return tw
}
