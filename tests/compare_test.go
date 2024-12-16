package sliceutil

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/devrob-go/sliceutil"
)

func TestCompareSlices(t *testing.T) {
	assert.True(t, sliceutil.CompareSlices([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.False(t, sliceutil.CompareSlices([]int{1, 2, 3}, []int{3, 2, 1}))
	assert.False(t, sliceutil.CompareSlices([]int{1, 2, 3}, []int{1, 2}))
	assert.True(t, sliceutil.CompareSlices([]string{"a", "b", "c"}, []string{"a", "b", "c"}))
}

func TestFindDifferences(t *testing.T) {
	assert.Equal(t, sliceutil.FindDifferences([]int{1, 2, 3}, []int{2, 3, 4}), []int{1, 4})
	assert.Equal(t, sliceutil.FindDifferences([]string{"a", "b"}, []string{"b", "c"}), []string{"a", "c"})
}

func TestMergeSlices(t *testing.T) {
	// Ascending order merge
	assert.Equal(t, sliceutil.MergeSlices([]int{1, 3, 5}, []int{2, 4, 6}, true, func(a, b int) bool {
		return a < b
	}), []int{1, 2, 3, 4, 5, 6})

	// Descending order merge
	assert.Equal(t, sliceutil.MergeSlices([]int{1, 3, 5}, []int{2, 4, 6}, false, func(a, b int) bool {
		return a < b
	}), []int{6, 5, 4, 3, 2, 1})
}

func TestMaxMinInt(t *testing.T) {
	assert.Equal(t, sliceutil.MaxInt([]int{1, 2, 3, 4}), 4)
	assert.Equal(t, sliceutil.MinInt([]int{1, 2, 3, 4}), 1)
}

func TestCompareSum(t *testing.T) {
	assert.Equal(t, sliceutil.CompareSum([]int{1, 2, 3}, []int{4, 5, 6}), "b is greater")
	assert.Equal(t, sliceutil.CompareSum([]int{1, 2}, []int{1, 2}), "both are equal")
	assert.Equal(t, sliceutil.CompareSum([]int{10, 20}, []int{5, 5}), "a is greater")
}
