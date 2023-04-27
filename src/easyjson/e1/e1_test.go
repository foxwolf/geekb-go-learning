package e1

import (
	"fmt"
	"testing"
)

type BasicInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

//easyjson:json 这个注释是必须要的，否则easyJson会忽略这个非struct类型
//easyjson:json
type BasicInfoList []BasicInfo

func TestEasyJson(t *testing.T) {
	e1 := BasicInfo{"Mike", 10}
	e2 := BasicInfo{"Rose", 11}
	eList := BasicInfoList{e1, e2}
	v, _ := eList.MarshalJSON()
	fmt.Println(string(v))
}