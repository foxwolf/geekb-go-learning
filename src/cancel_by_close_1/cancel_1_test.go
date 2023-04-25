package cancelbyclose_test

import (
	"context"
	"fmt"
	// "os"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done(): //接收取消通知
		return true
	default:
		return false
	}
}

func TestCancel1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background()) //根context,返回的是个function

	for i := 0; i < 5; i++ {
		// go func(i int, ctx context.Context) { //方法定义
		go func(i int) { //方法定义
			for {
				if isCancelled(ctx) {
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
		}(i) //方法调用
		// }(i, ctx) //方法调用
	}
	

	time.Sleep(time.Second * 4)
	cancel()
	time.Sleep(time.Second * 1)
}
