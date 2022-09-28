package collection

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTimingWheel(t *testing.T) {
	tw := NewTimingWheel(60, time.Millisecond*10)
	type arg struct {
		caseName string
		f        func(tw *TimingWheel)
	}
	cases := []arg{
		{
			caseName: "base",
			f: func(tw *TimingWheel) {

			},
		},
		{
			caseName: "sleep_2s",
			f: func(tw *TimingWheel) {
				time.Sleep(time.Second * 2)
			},
		},
		{
			caseName: "sleep_5.52s",
			f: func(tw *TimingWheel) {
				time.Sleep(time.Millisecond * 5520)
			},
		},
	}
	for _, myCase := range cases {
		myCase := myCase
		t.Run(myCase.caseName, func(t *testing.T) {
			if myCase.f != nil {
				myCase.f(tw)
			}
			start := time.Now()
			fmt.Printf("%s start: %s\n", myCase.caseName, start.Format("04:05.000"))
			tw.AddTask("1 sec", time.Second*1, func(ID string) {
				fmt.Printf("[case=%s]\tid=%s, executedAt: %s, cost: %d ms\n", myCase.caseName, ID, time.Now().Format("04:05.000"), time.Now().Sub(start).Milliseconds())
			})
			tw.AddTask("2 sec", time.Second*2, func(ID string) {
				fmt.Printf("[case=%s]\tid=%s, executedAt: %s, cost: %d ms\n", myCase.caseName, ID, time.Now().Format("04:05.000"), time.Now().Sub(start).Milliseconds())
			})
			tw.AddTask("20 sec", time.Second*20, func(ID string) {
				fmt.Printf("[case=%s]\tid=%s, executedAt: %s, cost: %d ms\n", myCase.caseName, ID, time.Now().Format("04:05.000"), time.Now().Sub(start).Milliseconds())
			})
			tw.AddTask("500 ms", time.Millisecond*500, func(ID string) {
				fmt.Printf("[case=%s]\tid=%s, executedAt: %s, cost: %d ms\n", myCase.caseName, ID, time.Now().Format("04:05.000"), time.Now().Sub(start).Milliseconds())
			})
			tw.AddTask("1 m", time.Minute*1, func(ID string) {
				fmt.Printf("[case=%s]\tid=%s, executedAt: %s, cost: %d ms\n", myCase.caseName, ID, time.Now().Format("04:05.000"), time.Now().Sub(start).Milliseconds())
			})
		})
	}

	t.Run("slot_mult_task", func(t *testing.T) {
		start := time.Now()
		tw.AddTask("1_600ms", time.Millisecond*600, func(ID string) {
			fmt.Printf("[case=%s]\tid=%s, executedAt: %s, cost: %d ms\n", "1_600ms", ID, time.Now().Format("04:05.000"), time.Now().Sub(start).Milliseconds())
		})
		tw.AddTask("1_1200ms", time.Millisecond*1200, func(ID string) {
			fmt.Printf("[case=%s]\tid=%s, executedAt: %s, cost: %d ms\n", "1_1200ms", ID, time.Now().Format("04:05.000"), time.Now().Sub(start).Milliseconds())
		})
		tw.AddTask("1_1800ms", time.Millisecond*1800, func(ID string) {
			fmt.Printf("[case=%s]\tid=%s, executedAt: %s, cost: %d ms\n", "1_1800ms", ID, time.Now().Format("04:05.000"), time.Now().Sub(start).Milliseconds())
		})
		tw.AddTask("1_2400ms", time.Millisecond*2400, func(ID string) {
			fmt.Printf("[case=%s]\tid=%s, executedAt: %s, cost: %d ms\n", "1_2400ms", ID, time.Now().Format("04:05.000"), time.Now().Sub(start).Milliseconds())
		})
	})
	go func() {
		// 2m 后关闭
		time.Sleep(time.Minute * 2)
		tw.Stop()
	}()
	tw.Wait()
}

func toArray(t *Task) []string {
	res := make([]string, 0)
	for t != nil {
		res = append(res, t.ID)
		t = t.Next
	}
	fmt.Printf("%v\n", res)
	return res
}
