package singleton_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {

}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		// singleInstance = new(Singleton)
		singleInstance = &Singleton{}
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func ()  {
			obj := GetSingletonObj()
			fmt.Printf("%v\n", unsafe.Pointer(obj))	
			wg.Done()
		}()	
	}

	wg.Wait()
}

func TestRunByOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			singleton := GetSingletonObj()
			fmt.Printf("%v %p \n", i, singleton)
		}(i)
	}
	wg.Wait()
}