package objpool

import (
	"errors"
	"fmt"
	"time"
	"unsafe"
)

type ReusableObj struct {
	id int  //结构体，用unsafe.Sizeof发现占用的大小为0，所以每次初始化的位移就是0，所以每个对象的地址都是相同。
					//而如果增加一个int成员的话，Size就会变成8，然后地址的位移也是8.
					//GO中为了优化如果结构体里没有字段只有方法，那么两个结构体的执行结果肯定相同，所以共用一个地址了。
}

type ObjPool struct {
	bufChan chan *ReusableObj //用于缓冲可重用对象
}

func NewObjPool(numOfObj int) *ObjPool {
	ObjPool := ObjPool{}
	ObjPool.bufChan = make(chan *ReusableObj, numOfObj)

	for i := 0; i < numOfObj; i++ {
		obj := &ReusableObj{}
		fmt.Printf("obj: %v\n", unsafe.Pointer(obj))
		ObjPool.bufChan <- obj
	}

	return &ObjPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("OverFlow")
	}
}
