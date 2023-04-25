package emptyinterface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//断言
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }

	// if i, ok := p.(string); ok {
	// 	fmt.Println("string", i)
	// 	return
	// }

	// fmt.Println("Unknow Type")
	switch v := p.(type) {
	case int:
		fmt.Printf("Integer, %T\n", v)
	case string:
		fmt.Printf("string %T\n", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	DoSomething(0.2)
}
