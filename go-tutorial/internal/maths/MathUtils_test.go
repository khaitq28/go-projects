package maths

import "testing"

func TestFactorial(t *testing.T) {
	tests := testData{
		{5, 120},
		{1, 1},
		{4, 24},
	}
	for _, test := range tests {
		result := Factorial(test.input)
		if result != test.output {
			t.Error(test.input, test.output, result)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{5, 7, 12},
		{12, 33, 45},
	}
	for _, test := range tests {
		val := Sum(test.a, test.b)
		if val != test.c {
			t.Error(test.a, test.b, test.c, val)
		}
	}
}

type testData []struct {
	input  int
	output int
}
