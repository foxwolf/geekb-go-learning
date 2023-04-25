package polymorphism_test

import (
	"fmt"
	"testing"
)

type Code string
type Program interface {
	WriteHelloWorld() Code
}

type goProgram struct {
}

type javaProgram struct {
}

func (p *goProgram) WriteHelloWorld() Code {
	return "fmt.println(\"Hello World!\")"
}

// func (p *javaProgram) WriteHelloWorld() Code {
// 	return "System.out.Println(\"Hello World!\")"
// }

func (p javaProgram) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Program) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goPro := new(goProgram)
	// javaPro := new(javaProgram)
	javaPro := &javaProgram{} //这里传值也是可以的。 javaPro := javaProgram{}

	writeFirstProgram(goPro)
	writeFirstProgram(javaPro)
}
