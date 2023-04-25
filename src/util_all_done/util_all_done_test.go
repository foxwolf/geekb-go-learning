package untilalldone_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numOfrunner := 10
	ch := make(chan string, numOfrunner)

	for i := 0; i < numOfrunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	findRet := ""
	for i := 0; i < numOfrunner; i++ {
		findRet += <-ch + "\n" //字符串拼接，并不高效
	}

	return findRet
}

func AllResponseWithRange() {
	numOfChannel := 10
	var wg sync.WaitGroup

	ch := make(chan string, numOfChannel)

	for i := 0; i < numOfChannel; i++ {
		wg.Add(1)	
		go func (i int)  {
			ch <- fmt.Sprintf("The result is from %d", i)	
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(ch)
	totalResult := ""

	for v := range ch {
		totalResult += v + "\n"	
	}

	fmt.Println(totalResult)
}

func WWTestDDD(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	AllResponseWithRange()
	t.Log("after:", runtime.NumGoroutine())
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	// time.Sleep(time.Second)
	t.Log("after:", runtime.NumGoroutine())
}
