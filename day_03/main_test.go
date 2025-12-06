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
