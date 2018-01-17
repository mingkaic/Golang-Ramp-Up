package sortstringarray

import (
	"testing"
	"strings"
)

type sample struct {
	input []string
	expected []string
}

func TestSortString(t *testing.T) {

	samples := []sample {
		{[]string {"d","b","a","h"}, []string {"a","b","d","h"}},
		{[]string {"z","a"}, []string {"a","z"}},
	}

	for _, s := range samples {

		result := SortString(s.input)

		for i :=0; i < len(s.input); i++ {

			if strings.Compare(result[i], s.expected[i]) != 0 {
				t.Fatalf("TestSortStringArray Failed: result = %v, expected = %v", result, s.expected)
			}
		}
	}
}