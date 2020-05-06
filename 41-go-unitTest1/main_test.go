package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if calculate(5) != 7 {
		t.Error("Expected 5 + 2 to equal 7")
	}
}

func TestTableCalculate(t *testing.T) {
	var listTest = []struct {
		input    int
		expected int
	}{
		{5, 7},
		{2, 4},
		{1000, 1002},
		{9999, 10001},
	}

	for _, test := range listTest {
		if result := calculate(test.input); result != test.expected {
			t.Error("Expected ", test.input, "+ 2 to equal ", test.expected)
		}
	}
}
