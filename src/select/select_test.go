package csp_test

import (
	"fmt"
	"testing"
	"time"
)

// func TestService(t *testing.T) {
// 	fmt.Println(service())
// 	otherTask()
// }

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

// func otherTask() {
// 	fmt.Println("working on something else")
// 	time.Sleep(time.Millisecond * 100)
// 	fmt.Println("Task is done.")
// }

func AsyncService() chan string {
	// retCh := make(chan string)
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()

	return retCh
}

// func TestAsyncService(t *testing.T) {
// 	retCh := AsyncService()
// 	otherTask()
// 	fmt.Println(<-retCh)
// 	time.Sleep(time.Second * 1)
// }

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("Time out")
	}
	time.Sleep(time.Second * 3)
}
