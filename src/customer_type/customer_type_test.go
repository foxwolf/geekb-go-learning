package customertype_test

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(int)int

// 装饰者模式
func timeSpent(funp IntConv) func(int) int {
	return func(xx int) int {
		start := time.Now()
		ret := funp(xx)
		end := time.Now()

		fmt.Printf("time spent: %v\n", end.Sub(start))
		// fmt.Printf("time spent: %v\n", time.Since(start).Milliseconds())

		return ret
	}
}

func slowFun(ms int) int {
	// time.Sleep(time.Millisecond * 1)
	time.Sleep(time.Millisecond * time.Duration(ms))
	return ms
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	// tsSF(10)
	t.Log(tsSF(10))
}