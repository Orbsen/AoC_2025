package main

import "testing"

func TestDialRight(t *testing.T) {
	tests := []struct {
		testCase string
		dialPos  int
		amount   int
		expected int
	}{
		{"normal addition", 50, 10, 60},
		{"landing on 0", 90, 10, 0},
		{"roll over by 1", 90, 20, 10},
		{"roll over multiple times", 90, 120, 10},
	}

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			got := dialRight(tt.dialPos, tt.amount)
			if got != tt.expected {
				t.Errorf("DialRight(%d, %d) = %d; want %d", tt.dialPos, tt.amount, got, tt.expected)
			}
		})
	}
}

func TestDialLeft(t *testing.T) {
	tests := []struct {
		testCase string
		dialPos  int
		amount   int
		expected int
	}{
		{"normal subtraction", 50, 10, 40},
		{"landing on 0", 20, 20, 0},
		{"roll over by 1", 20, 30, 90},
		{"roll over multiple times", 20, 130, 90},
	}

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			got := dialLeft(tt.dialPos, tt.amount)
			if got != tt.expected {
				t.Errorf("DialLeft(%d, %d) = %d; want %d", tt.dialPos, tt.amount, got, tt.expected)
			}
		})
	}
}
