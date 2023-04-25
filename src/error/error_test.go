package error_test

import (
	"errors"
	"fmt"
	"testing"
)

var ErrfoLessThanTwoError = errors.New("n should be not less than 2")
var ErrfoLargerThanHundredError = errors.New("n shout be no larger than 100")

//及早失败，避免嵌套
func GetFibonacci(n int) ([]int, error) {
	// if n < 2 || n > 100 {
	// 	return nil, errors.New("n should be in [2,100]")
	// }

	if n < 2 {
		return nil, ErrfoLessThanTwoError
	}

	if n > 100 {
		return nil, ErrfoLargerThanHundredError
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1); err != nil {
		if err == ErrfoLessThanTwoError {
			fmt.Println("It is less 2")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestNewEqual(t *testing.T) {
	if errors.New("abc") == errors.New("abc") {
		t.Error("new(\"abc\")==New(\"abc\")")
	}

	if errors.New("abc") == errors.New("xyz") {
		t.Error("new(\"abc\") == new(\"xyz\")")
	}

	err := errors.New("jkl")
	if err != err {
		t.Error("err != err")
	}
}
