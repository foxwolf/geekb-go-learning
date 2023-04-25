package main

import (
	"fmt"
	"log"
	"runtime"
	// "time"
)

func TT() {
	log.Println("0")
}
func main() {
	fmt.Printf("cpu: %d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go func() {
		// TT()
		for {
			log.Println("0")
			// time.Sleep(time.Second * 1)
		}
	}()
	for {

		log.Println("1")
		// time.Sleep(time.Second * 1)
	}
}
