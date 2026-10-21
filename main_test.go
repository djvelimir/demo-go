package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{5, 120},
		{9, 362880},
	}

	for _, tt := range tests {
		got := factorial(tt.n)
		if got != tt.want {
			t.Errorf("factorial(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{5, 5},
		{9, 34},
	}

	for _, tt := range tests {
		got := fibonacci(tt.n)
		if got != tt.want {
			t.Errorf("fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
