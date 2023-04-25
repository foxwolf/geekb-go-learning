package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"

	parts := strings.Split(s, ",")
	for _, v := range parts {
		t.Log(v)
	}

	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(6)

	t.Logf("s:%s, %T", s, s)
	if v, err := strconv.Atoi("12"); err == nil {
		t.Log(10 + v)
	} else {
		t.Log("error")
	}
}
