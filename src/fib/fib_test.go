package fib

import "testing"

func TestFib(t *testing.T) {
	x := 1
	y := 1

	t.Log(x)
	for i := 1; i < 11; i++ {
		x, y = y, x+y
		t.Log(" ", x)
	}
}
func TestSwp(t *testing.T) {
	a := -1
	b := -8

	// a = a + b
	// b = a - b
	// a = a - b

	a, b = b, a
	t.Log(a, b)
}
