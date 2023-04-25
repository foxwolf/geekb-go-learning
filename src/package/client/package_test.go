package client_test

import (
	"testing"
	"github.com/foxwolf/geekb-go-learning/package/series"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
	// t.Log(series.squart(3))
}
