package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 3 {
		t.Error("expected 2+2 equals 4")
	}
}

func TestTableCalculates(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{9999, 10001},
	}
	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} input,{} expected recived {}", test.input, test.expected, output)
		}
	}

}
