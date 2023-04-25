package objpool

import (
	"fmt"
	"time"
	"unsafe"
	"testing"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(5)

	// go func() {
	// 	for i := 0; i < 1; i++ {
	// 		if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	// 			t.Error(err)
	// 		}
	// 	}
	// }()

	// time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%v\n", unsafe.Pointer(v))
			// if err := pool.ReleaseObj(v); err != nil {
			// 	t.Error(err)
			// }
		}
	}
}
