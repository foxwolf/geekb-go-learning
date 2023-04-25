package flexiblereflect

import (
	"errors"
	"fmt"

	// "os"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}

	// t.Log(a == b)
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	// t.Log(s1 == s2)
	t.Log("s1==s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1==s3?", reflect.DeepEqual(s1, s3))
}

func fillBySettings(st interface{}, setting map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer")
	}

	// Elem()获取指针指向的值
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type")
	}

	if setting == nil {
		return errors.New("setting is nil.")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range setting {
		if field, ok = (reflect.ValueOf(st).Elem().Type().FieldByName(k)); !ok {
			continue
		}

		// fmt.Println("--", field.Type, "---", reflect.TypeOf(v))
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			// fmt.Println(vstr)
			vstr = vstr.Elem()
			// fmt.Println(st)
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	// if v, ok := reflect.ValueOf(st).Elem().Type().FieldByName("name"); !ok {
	// 	fmt.Println("--", v)
	// }
	// -- {  <nil>  0 [] false}

	// fmt.Println("--", (reflect.ValueOf(st).Type()))
	// -- *flexiblereflect.Employee
	// -- *flexiblereflect.Customer
	// fmt.Println("--", (reflect.ValueOf(st)))
	// -- &{ Mike 30}
	// fmt.Println("--", (reflect.ValueOf(st)).Elem())
	// fmt.Println("**", (reflect.ValueOf(st)).Elem().Type())
	// -- { Mike 30}
	// ** flexiblereflect.Employee
	// -- { Mike 30}
	// ** flexiblereflect.Customer
	return nil
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

type Customer struct {
	CookieID string
	name     string
	Age      int
}

func TestFillNameAndAge(t *testing.T) {
	setting := map[string]interface{}{"Name": "Mike", "Age": 30}
	e := Employee{"1001", "jack", 22}

	fmt.Println(reflect.ValueOf(&e).Elem().Type().FieldByName("Name"))
	fmt.Println(reflect.ValueOf(&e).Elem().FieldByName("Name"))
	fmt.Println(reflect.TypeOf(&e).Elem().FieldByName("Name"))
// {Name  string format:"normal" 16 [1] false} true
// jack
// {Name  string format:"normal" 16 [1] false} true

	if err := fillBySettings(&e, setting); err != nil {
		t.Fatal(err)
	}

	t.Log(e)

	c := new(Customer)
	if err := fillBySettings(c, setting); err != nil {
		t.Fatal(err)
	}

	t.Log(*c)
}
