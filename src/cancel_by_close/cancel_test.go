package cancelbyclose_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct {}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{} //发的空结构，就是个空消息，没有任何意义。
}

func cancel_2(cancelChan chan struct{}) {
	// cancelChan <- struct{}{}
	close(cancelChan)
}

func TestCancel1(t *testing.T) {
	cancelChan := make(chan struct{},1)

	for i := 0; i < 5; i++ {
		go func (i int, cancelCh chan struct{})  {
			for {
				if isCancelled(cancelCh)	 {
					break
				}
				fmt.Printf("go routine: %v\n", i)
				// if _, ok := <- cancelCh; ok {
				// 	break
				// }
				fmt.Printf("go routine: %v\n", i)
				time.Sleep(time.Millisecond * 500)
			}	
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}

	time.Sleep(time.Second * 1)
	// cancel_1(cancelChan)
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{},1)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func (i int, cancelCh chan struct{})  {
			for {
				if isCancelled(cancelCh)	 {
					break
				}
				time.Sleep(time.Millisecond * 5000)
			}	
			fmt.Println(i, "Cancelled")
			wg.Done()
		}(i, cancelChan)
	}

	// cancel_1(cancelChan)
	cancel_2(cancelChan)
	wg.Wait()
}