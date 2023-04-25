package map_test

import (
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 4: 9}

	t.Log(m1[4], "len: ", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Log(m2[2], "len: ", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

//map不存在时会初始化为0值
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}

	m1[2] = 0
	m1[3] = 0
	if v, err := m1[3]; err {
		t.Log("key 3 is existing...", v)
	} else {
		t.Log("key 3 is not existing...", v)
	}
}

func TestTraveMap(t *testing.T) {
	m1 := map[int]int{1:2, 2:4, 9:9}

	for i, v:= range m1 {
		t.Log("key",i,"value=",v)
	}
}
