package utilanyonereply_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfrunner := 10
	ch := make(chan string, numOfrunner)
	// ch := make(chan string) //注意这块会导致协程泄漏

	for i := 0; i < numOfrunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	xx := <-ch
	wg.Done()
	return xx
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	wg.Add(1)
	t.Log(FirstResponse())
	wg.Wait()
	time.Sleep(time.Second)
	t.Log("after:", runtime.NumGoroutine())
}


