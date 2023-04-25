package interface_test

import (
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

//方法签名 duck type
func (go1 *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"hello world\")"
}

func TestClient(t *testing.T) {
	// var p Programmer
	p := new(GoProgrammer)

	t.Log(p.WriteHelloWorld())
}
