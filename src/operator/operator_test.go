package operator

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 3, 2, 4}
	d := [...]int{1, 3, 3, 4}

	t.Log(a == b, a == c, b == c, c == d)
}

const (
	readAble = 1 << iota
	writeAble
	execAble
)

func TestLogic(t *testing.T) {
	bitMask := 0x07

	t.Log(bitMask&execAble == execAble) //true

	bitMask = bitMask &^ execAble
	t.Log(bitMask&execAble == execAble) //false
}

//数组赋值与切片赋值不同
func TestArraySlice(t *testing.T) {
	a := [...]int{1,2,3,4}
	s := []int{1,2,3,4}

	b := a
	b[2] = 10
	t.Log(a)
	t.Log(b)

	b1 := s
	b1[2] = 10
	t.Log(s)
	t.Log(b1)
}
