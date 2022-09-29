package rescue

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRunSafe(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go RunSafe(func() {
			defer wg.Done()
			if i%2 == 0 {
				panic(fmt.Sprintf("%d panic", i))
			}
		})
	}
	wg.Wait()
	time.Sleep(time.Second * 2)
}
