package switchStatement

import (
"testing"
)


type sample struct {
	input []int
	expected [] int
}


func TestArrayAdd(t *testing.T) {

	samples := []sample {
		{[]int {1,2,3,4}, []int {2,3,4,5}},
		{[]int {1}, []int {2}},
		{[]int {}, nil},
	}


	for _, s := range samples {

		result := ArrayAdd(s.input)

		for i :=0; i < len(s.input); i++ {

			if result[i] != s.expected[i] {
				t.Fatalf("TestArrayAdd Failed: result = %v, expected = %v", result, s.expected)
			}
		}


	}


}
