package utils

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "no duplicates",
			input:    []string{"apple", "banana", "orange"},
			expected: []string{"apple", "banana", "orange"},
		},
		{
			name:     "with duplicates",
			input:    []string{"apple", "banana", "apple", "orange", "banana"},
			expected: []string{"apple", "banana", "orange"},
		},
		{
			name:     "all duplicates",
			input:    []string{"apple", "apple", "apple"},
			expected: []string{"apple"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RemoveDuplicates(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("RemoveDuplicates(%v) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

// BenchmarkRemoveDuplicatesString benchmarks RemoveDuplicates for strings.
func BenchmarkRemoveDuplicatesString(b *testing.B) {
	input := []string{"apple", "banana", "apple", "orange", "banana", "grape", "apple", "kiwi", "banana", "orange"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicates(input)
	}
}

// BenchmarkRemoveDuplicatesInt benchmarks RemoveDuplicates for integers.
func BenchmarkRemoveDuplicatesInt(b *testing.B) {
	input := []int{1, 2, 3, 1, 2, 4, 5, 1, 2, 3, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicates(input)
	}
}
