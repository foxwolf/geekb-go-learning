package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	// p *Pet
	Pet //匿名嵌套类型
}

func (d *Dog) Speak() {
	// d.p.Speak()
	fmt.Print("wang!")
}

// func (d *Dog) SpeakTo(host string) {
// 	// d.p.SpeakTo(host)
// 	d.Speak()
// 	fmt.Println(" ", host)
// }

func TestDog(t *testing.T) {	
	dog := new(Dog)
	// dog.Speak()
	dog.SpeakTo("chao")
}