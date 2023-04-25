package cancelwithctx_test

import (
	"fmt"
	"testing"
)

func WithContext(ctx int) func () {
	return func()	 {
		fmt.Println("The value of ctx is ", ctx)
	}
}

func TestWithContext(t *testing.T) {
	cancel := WithContext(10)	
	cancel()
}