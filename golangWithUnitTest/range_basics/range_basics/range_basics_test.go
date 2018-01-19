package range_basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeBasicsNoNumber(t *testing.T) {
	nums := []int{}
	assert.Equal(t, 0, RangeBasics(nums), "Wrong sum.")
}

func TestRangeBasicsOneNumber(t *testing.T) {
	nums := []int{1}
	assert.Equal(t, 1, RangeBasics(nums), "Wrong sum.")
}

func TestRangeBasics(t *testing.T) {
	nums := []int{1,2,3,4}
	assert.Equal(t, 10, RangeBasics(nums), "Wrong sum.")
}

