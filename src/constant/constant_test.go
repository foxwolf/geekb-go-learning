package constant

import (
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Unkown = 1 << iota
	On
	Off
)

func TestConst(t *testing.T) {
	t.Log(Monday, Wednesday)
}

func TestConst1(t *testing.T) {
	a := 0x55 //0101 0101

	t.Log(a & Unkown) //1
	t.Log(a & On)     //0
	t.Log(a & Off)    //4
	t.Log(a&Unkown == Unkown, a&On==On, a&Off==Off) //true false true
}
