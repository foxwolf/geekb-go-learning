package csp_test

import (
	"fmt"
	"testing"
	"time"
)

type Custome struct {
	Id   string
	Name string
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func AsyncService() chan string {
	retCh := make(chan string)
	// retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()

	return retCh
}

func service1(p *Custome) *Custome {
	p.Id = "123"
	p.Name = "apple"
	return p
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	// time.Sleep(time.Second * 1)
}

func AsyncService1(p *Custome) chan Custome {
	retCh := make(chan Custome)
	// retCh := make(chan string, 1)

	go func() {
		ret := service1(p)
		fmt.Println("returned result.")
		retCh <- *ret
		fmt.Println("service exited.")
	}()

	return retCh
}

func TestAsyncService1(t *testing.T) {
	// p := &Custome{}
	p := new(Custome)

	retCh := AsyncService1(p)
	otherTask()
	fmt.Printf("%+v\n", <-retCh)
	// time.Sleep(time.Second * 1)
}