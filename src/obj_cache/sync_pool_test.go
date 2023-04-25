package objcache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v, ok := pool.Get().(int) //断言 (int)
	fmt.Println(v, ok)
	pool.Put(3)
	runtime.GC()
	v1, ok :=pool.Get().(int)
	fmt.Println(v1, ok)
}

func ATestSyncPoolInMultiGrouting(t *testing.T) {
	pool := &sync.Pool{
		New: func() any {
			fmt.Println("Create a new Object")	
			return 10
		},
	}

	pool.Put(100)
	pool.Put(101)
	pool.Put(102)
	pool.Put(103)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Printf("%d %d\n", id, pool.Get())	
			wg.Done()
		}(i)
	}

	wg.Wait()
}