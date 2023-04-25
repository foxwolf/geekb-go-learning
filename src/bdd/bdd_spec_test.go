package bdd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey" //.表示将import进来的package的方法是在当前名字空间的，可以直接使用。
)

func TestSpec(t *testing.T) {
	Convey("Given 2 even numbers", t, func ()  {
		a := 3
		b := 2	

		Convey("when add the two numbers", func ()  {
			c := a + b

			Convey("Then the result is still even", func ()  {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}