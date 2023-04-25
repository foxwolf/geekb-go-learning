package unsafeprogramming

import (
	. "fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

/*注意使用不安全模块编程虽然可以强制类型转换，但数据是不对的，危险!!!*/
func TestUnSafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))

	Println(unsafe.Pointer(&i))
	Println(f)
	// 0xc00001c168
	// 5e-323
}

type MyInt int

func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	Println(b)
}

// 原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer

	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup
	writeDataFn()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
		}()

		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
