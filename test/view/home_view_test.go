package view_test

import (
	"prod_tracker/view"
	"testing"
)

func TestTranslateSeconds(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "00:00:00"},
		{59, "00:00:59"},
		{60, "00:01:00"},
		{3600, "01:00:00"},
		{3661, "01:01:01"},
		{86399, "23:59:59"}, // Edge case: last second of the day
	}

	for _, test := range tests {
		result := view.TranslateSeconds(test.input)
		if result != test.expected {
			t.Errorf("TranslateSeconds(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestGetActivityNames(t *testing.T) {
	tests := []struct {
		input    [][]string
		expected []string
	}{
		{[][]string{}, []string{}},                          // Empty input
		{[][]string{{"1", "Running"}}, []string{"Running"}}, // Single activity
		{[][]string{{"1", "Running"}, {"2", "Swimming"}, {"3", "Cycling"}}, []string{"Running", "Swimming", "Cycling"}}, // Multiple activities
	}

	for _, test := range tests {
		result := view.GetActivityNames(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("GetActivityNames(%v) returned %v; want %v", test.input, result, test.expected)
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("GetActivityNames(%v) = %v; want %v", test.input, result, test.expected)
			}
		}
	}
}
