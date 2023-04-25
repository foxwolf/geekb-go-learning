package client_test

import (
	"geekb-go-learn/src/package/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
	// t.Log(series.squart(3))
}
