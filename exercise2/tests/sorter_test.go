package tests

import (
	"reflect"
	"testing"

	"github.com/diegovanne/go23/exercise2/internal/sorter"
)

func TestSortMixedArray(t *testing.T) {
	tests := []struct {
		name     string
		input    sorter.MixedSorter
		expected sorter.MixedSorter
	}{
		// Test cases with mixed types
		{
			"case1",
			sorter.MixedSorter{10, 2.5, "apple", 5, 1.5},
			sorter.MixedSorter{1.5, 2.5, 5, 10, "apple"},
		},
		{
			"case2",
			sorter.MixedSorter{-10, "banana", 7.3, "orange", 0},
			sorter.MixedSorter{-10, 0, 7.3, "banana", "orange"},
		},
		{
			"case3",
			sorter.MixedSorter{1, "cat", 3.14, 0, -5.5},
			sorter.MixedSorter{-5.5, 0, 1, 3.14, "cat"},
		},
		{
			"case4",
			sorter.MixedSorter{},
			sorter.MixedSorter{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy the input to a new slice as it will be modified by SortMixedArray
			inputCopy := make(sorter.MixedSorter, len(tt.input))
			copy(inputCopy, tt.input)

			sorter.SortMixedArray(inputCopy)

			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("SortMixedArray() got = %v, expected %v", inputCopy, tt.expected)
			}
		})
	}
}

func TestSortMixedNumber(t *testing.T) {
	tests := []struct {
		name     string
		ints     []int
		floats   []float64
		expected sorter.MixedSorter
	}{
		// Test cases with mixed types
		{
			"case1",
			[]int{10, -5, 2},
			[]float64{3.14, 1.5, 2.75},
			sorter.MixedSorter{-5, 1.5, 2, 2.75, 3.14, 10},
		},
		{
			"case2",
			[]int{-1, -2, -3},
			[]float64{1.1, 2.2, 3.3},
			sorter.MixedSorter{-3, -2, -1, 1.1, 2.2, 3.3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sorter.SortMixedNumber(tt.ints, tt.floats)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SortMixedNumber() got = %v, expected %v", got, tt.expected)
			}
		})
	}
}
