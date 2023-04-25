package channelclose_test

import (
	"fmt"
	"sync"
	"testing"
	// "time"
	// "time"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// time.Sleep(time.Millisecond * 100)
		for i := 0; i < 15; i++ {
			ch <- i
		}
		close(ch)
		fmt.Println("Close Channel")
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup, num int) {
	go func() {
		for {
			// time.Sleep(time.Millisecond * 100)
			if value, ok := <-ch; ok {
				fmt.Printf("%v,%v\n", num,value)
				// if value == 4 {
				// 	break
				// }
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan int)
	// ch := make(chan int, 10)
	wg.Add(1)
	dataReceiver(ch, &wg, 1)
	wg.Add(1)
	dataProducer(ch, &wg)
	// wg.Add(1)
	// dataReceiver(ch, &wg, 2)
	// wg.Add(1)
	// dataReceiver(ch, &wg, 3)
	wg.Wait()

}
