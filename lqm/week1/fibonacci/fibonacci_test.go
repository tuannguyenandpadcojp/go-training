package main

import (
	"fmt"
	"testing"
)

func Test_fib(t *testing.T) {
	tests := [][2]int{
		{-1, -1},
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}
	for _, tt := range tests {
		tc := fmt.Sprintf("F(%d) = %d", tt[0], tt[1])
		t.Run(tc, func(t *testing.T) {
			result := fibonacci(tt[0])
			expected := tt[1]
			assertCorrectResult(t, result, expected)
		})
	}
}

func assertCorrectResult(t testing.TB, result, expected int) {
	t.Helper()
	if result != expected {
		t.Errorf("expected '%d' but got result '%d'", expected, result)
	}
}
