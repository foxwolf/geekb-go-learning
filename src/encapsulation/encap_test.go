package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s - Name:%s - Age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee) String1() string {
	fmt.Printf("Address1 is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("%+v; %T\n", e, e)

	e = Employee{Id: "1", Name: "Mike", Age: 20}
	// t.Log(e)
	fmt.Printf("%+v; %T\n", e, e)

	//返回指针
	e1 := new(Employee)
	e1.Id = "2"
	e1.Name = "Rose"
	e1.Age = 20
	fmt.Printf("%+v; %T\n", e1, e1) //&{Id:2 Name:Rose Age:20}

	fmt.Printf("point is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
	t.Log(e.String1())
}

