package array_test

import "testing"

func TestArray(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{4, 5, 678, 8, 3, 1}
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr2, len(arr2))

	arr3 := arr1[1:]
	t.Log("arr3", arr3)
}

func TestArrayTravel(t *testing.T) {
	a := [...]int{1,2,3,4,5}

	for _, e := range a {
		t.Log(e)
	}
}