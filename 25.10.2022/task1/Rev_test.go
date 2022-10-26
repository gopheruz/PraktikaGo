package test

import "testing"

func TestReverse(t *testing.T) {
	TestestCases := []struct{
		input int
		output int
	}{
		{12,21},
		{22, 22},
		{243, 342},
	}
	for _, testCase := range TestestCases{
		if testCase.output != Reverse(testCase.input){
			t.Error("Error reversing number")
		}
	}
}
