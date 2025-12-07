package main

import "testing"

func TestGetHighestValueAndIndex(t *testing.T) {
	tests := []struct {
		testCase      string
		expectedValue string
		expectedIndex int
		withLastValue bool
		input         string
	}{
		{
			"default",
			"9",
			3,
			true,
			"345912",
		},
		{
			"without lastValue",
			"4",
			1,
			false,
			"348",
		},
	}
	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			gotValue, gotIndex := getHighestValueAndIndex(tt.input, tt.withLastValue)
			if gotValue != tt.expectedValue || gotIndex != tt.expectedIndex {
				t.Errorf("got Value: %s excpetedValue: %s got Index: %d excpectedIndex: %d", gotValue, tt.expectedValue, gotIndex, tt.expectedIndex)
			}
		})
	}
}

func TestGetHighestJoltage(t *testing.T) {
	tests := []struct {
		testCase string
		input    string
		length   int
		expected int
	}{
		{"lowest Numbers are last", "8765432", 5, 87654},
		{"lowest Numbers are first", "1234567", 5, 34567},
		{"lowest Numbers are in between", "3526765", 5, 56765},
	}

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			got := getHighestJoltage(tt.input, tt.length)
			if got != tt.expected {
				t.Errorf("got %d excpeted: %d ", got, tt.expected)
			}
		})
	}
}
