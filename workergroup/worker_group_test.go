package workergroup

import (
	"fmt"
	"testing"
	"time"
)

func TestTG(t *testing.T) {
	tg := New(10)
	tpl := "04:05.000"
	for i := 0; i < 100; i++ {
		i := i
		tg.Run(func() {
			start := time.Now()
			defer func() {
				fmt.Printf("[%d] start: %s, end: %s\n", i, start.Format(tpl), time.Now().Format(tpl))
			}()
			time.Sleep(time.Second * 2)
		})
	}
	tg.Wait()
}
