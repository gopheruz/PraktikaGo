package test

import "testing"

// func TestHelloWorld(t *testing.T) {
// 	result := Hello()
// 	check := "Hello World!"
// 	if check != result{
// 		t.Error("Error printing hello world")
// 	}
// }

func TestSalomlash(t *testing.T) {
	testCases := []struct{
		input string
		lang string
		output1 string
	}{
		{"Nurmuhammad", "EN", "Hello Nurmuhammad"},
		{"Nurmuhammad", "RU", "Privet Nurmuhammad"},
		{"Nurmuhammad", "UZB", "Salom Nurmuhammad"},
	}
	for _, testCase := range testCases{
		if testCase.output1 != Hello(testCase.input, testCase.lang){
			t.Error("Error program fied", testCase)
		}
	}
}
