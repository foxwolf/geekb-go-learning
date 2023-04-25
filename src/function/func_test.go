package function_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValue() (int, int) {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(10), rand.Intn(20)
}

// 装饰者模式
func timeSpent(funp func(int) int) func(int) int {
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
	a, b := returnMultiValue()
	t.Log(a, b)
	tsSF := timeSpent(slowFun)
	tsSF(10)
	// t.Log(tsSF(10))
}

func sum(ops ...int) int {
	sum := 0

	for _, op := range ops {
		sum += op
	}

	return sum
}
func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
	t.Log(sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}
func TestDefer(t *testing.T) {
	// defer func()  {
	// 	t.Log("Clear resources")	
	// }()

	defer Clear()
	t.Log("Started")
	// panic("fatal error")
	t.Log("~~~~end~~~~~")
}