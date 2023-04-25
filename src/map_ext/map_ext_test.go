package map_ext_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}

	var bb bool

	t.Log(bb)

	m[0] = func(op int) int { return op }
	m[1] = func(op int) int { return op * op }
	m[2] = func(op int) int { return op * op * op }

	t.Log(m[0](2), m[1](2), m[2](2))
}

func TestMapForSet(t *testing.T) {
	m := map[int]bool{}

	m[1] = true
	n := 1
	if m[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}

	m[4] = true
	t.Logf("len: %d", len(m))

	delete(m, 1)
	n = 1
	if m[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
