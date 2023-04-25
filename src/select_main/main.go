package main

import (
	"fmt"
	"time"
)

func g1(c1 chan int) {
	c1 <- 21
}

func g2(c1 chan int) {
	c1 <- 22
}

func main1() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go g1(ch1)
	go g2(ch2)

	select {
	case v1 := <-ch1:
		fmt.Println("Got ", v1)
	case v2 := <-ch2:
		fmt.Println("Got ", v2)
	default:
		fmt.Println("Default")
	}
}

func main() {
	ch := make(chan int, 1)

	go func() {
		time.Sleep(time.Second * 8)
		ch <- 1
	}()

	select {
	case ch := <-ch:
		fmt.Println(ch)
	case <-time.After(time.Second * 10):
		fmt.Println("time out")
		// default:
		// 	fmt.Println("Default")
	}
}
