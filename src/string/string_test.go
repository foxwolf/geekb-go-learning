package string_test

import "testing"

func TestString(t *testing.T) {
	var str string
	t.Log(str)

	str = "hello"
	t.Log(str)
	t.Log("长度:", len(str)) //5
	// str[1] = 'Z'
	// t.Log(str)
	// t.Log("长度:", len(str))

	str = "你好"
	t.Log(str)
	t.Logf("长度: %d", len(str)) //6
	t.Logf("UTF8: %x", str) //e4 bd a0 e5 a5 bd

	// str = "\xE4\xBD\xA0" //你
	str = "\x4e\x2d" //你
	t.Log(str)
	t.Log("长度:", len(str)) //3

	str = "\xE4\xB8\xA5\xAA"
	t.Log(str)
	t.Log("长度:", len(str)) //4

	str = "中"
	t.Log(str)
	t.Log("长度:", len(str)) //3

	c := []rune(str)
	t.Log(len(c)) //1
	t.Logf("中 unicode %x", c[0]) //4e2d
	t.Logf("中 UTF8 %x", str) //e4 b8 ad
}

func TestStringToRun(t *testing.T) {
	s := "中华人民共和国"

	for _, v := range s {
		t.Logf("%[1]c %[1]d %[1]x", v)
	}
}