package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)

	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int64, reflect.Int32:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func iTestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)  //Float
	CheckType(&f) //Unknown *float64
}

func iTestTypeAndValue(t *testing.T) {
	var f int64 = 10

	t.Log(reflect.TypeOf(f), reflect.ValueOf(f)) //int64 10
	t.Log(reflect.ValueOf(f).Type())             //int64
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := Employee{"1", "Mike", 30}

	//按名字获取成员
	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(e).FieldByName("Name"); !ok {
		t.Error("Failed to get `Name` field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	t.Log("initialize Age:", e)
	reflect.ValueOf(&e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Update Age:", e)
}
