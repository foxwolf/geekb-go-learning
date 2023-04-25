package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for index, v := range inputs {
		ret := square(v)
		if ret != expected[index] {
			t.Errorf("input is %d, the expected is %d, the actual %d",
				v, expected[index], ret)
		}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("start")
	t.Error("Error")
	fmt.Println("end")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("start")
	t.Fatal("Error")
	fmt.Println("end")
}

func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for index, v := range inputs {
		ret := square(v)
		assert.Equal(t, expected[index], ret, "The Value should be the same")
		// if ret != expected[index] {
		// 	t.Errorf("input is %d, the expected is %d, the actual %d",
		// 		v, expected[index], ret)
		// }
	}
}