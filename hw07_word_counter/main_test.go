package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input: "Hello, world!",
			expected: map[string]int{
				"hello": 1,
				"world": 1,
			},
		},
		{
			input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Lorem ipsum.",
			expected: map[string]int{
				"lorem":       2,
				"ipsum":       2,
				"dolor":       1,
				"sit":         1,
				"amet":        1,
				"consectetur": 1,
				"adipiscing":  1,
				"elit":        1,
			},
		},
		{
			input: "One, two, three. Three, two, one.",
			expected: map[string]int{
				"one":   2,
				"two":   2,
				"three": 2,
			},
		},
		{
			input:    "  ",
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		result := countWords(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Для входной строки '%s' ожидался результат %v, но получен %v", test.input, test.expected, result)
		}
	}
}
