package main

import (
	"testing"
)

func TestThercle(t *testing.T) {
	testCases := []struct {
		a      int
		b      int
		c      int
		output bool
	}{
		{12, 10, 56, false},
		{6, 8, 10, true},
		{3, 4, 5, true},
	}
	for _, value := range testCases {
		if value.output != Thercle(value.a, value.b, value.c) {
			t.Error("Filed")
		}
	}
}
