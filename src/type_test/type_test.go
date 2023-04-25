package typetest

import (
	"math"
	"testing"
)

type MyInt int //类型再定义
// type MyInt = int

func TestType(t *testing.T) {
	var a, a1 int = 1, 2
	var b int64
	var c MyInt

	b = int64(a)
	a1 = a
	c = MyInt(a) //类型再这义支持强转
	// c = a

	t.Log(a, b)
	t.Log(a1, c)
	t.Log(math.MaxInt32)
}

// 不支持指针运算
func TestPoint(t *testing.T) {
	a := 10
	aPtr := &a

	*aPtr = 12
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string

	t.Log("*" + s + "*")
	t.Log(len(s))

	if s == "" {
	// if s == nil { //error
		t.Log("is nil")
	}

}
