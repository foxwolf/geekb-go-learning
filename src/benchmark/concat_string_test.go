package benchmark

import (
	"bytes"
	"fmt"
	"testing"

	_ "github.com/stretchr/testify/assert"
)

func TestConcatStringByAdd(t *testing.T) {
	// assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}

	ret := ""
	for i := 0; i < 5; i++ {
		for _, elem := range elems {
			ret += elem
		}
		fmt.Println(ret)
	}
	// assert.Equal("12345", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	// assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	var buf bytes.Buffer

	for i := 0; i < 5; i++ {
		for _, elem := range elems {
			buf.WriteString(elem)
		}
		fmt.Println(buf)
	}

	// assert.Equal("12345", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
