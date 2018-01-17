package sortNumberArray

import (
	"testing"
)

type sample struct {
	input []int
	expected []int
}

func TestSortArray(t *testing.T) {

	samples := []sample {
		{[]int {1,2,3,4}, []int {1,2,3,4}},
		{[]int {2,1}, []int {1,2}},
	}

	for _, s := range samples {

		result := SortArray(s.input)

		for i :=0; i < len(s.input); i++ {

			if result[i] != s.expected[i] {
				t.Fatalf("TestArraySort Failed: result = %v, expected = %v", result, s.expected)
			}
		}
	}
}
