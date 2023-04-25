package main

import (
	"fmt"
	_"runtime"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	// runtime.GOMAXPROCS(2)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Millisecond*500)
}
