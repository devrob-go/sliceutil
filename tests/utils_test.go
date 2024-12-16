package sliceutil

import (
	"testing"

	"github.com/devrob-go/sliceutil"
	"github.com/stretchr/testify/assert"
)

func TestCompareSlices(t *testing.T) {
	assert.True(t, sliceutil.CompareSlices([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.False(t, sliceutil.CompareSlices([]int{1, 2, 3}, []int{3, 2, 1}))
	assert.False(t, sliceutil.CompareSlices([]int{1, 2, 3}, []int{1, 2}))
	assert.True(t, sliceutil.CompareSlices([]string{"a", "b", "c"}, []string{"a", "b", "c"}))
}

func TestFindDifferences(t *testing.T) {
	t.Run("Test FindDifferences with int slices", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		b := []int{3, 4, 5, 6}
		expected := []int{1, 2, 5, 6}
		result := sliceutil.FindDifferences(a, b)

		// Use ElementsMatch to ignore order
		assert.ElementsMatch(t, expected, result, "The unique elements should match regardless of order")
	})

	t.Run("Test FindDifferences with string slices", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		b := []string{"banana", "date", "cherry"}
		expected := []string{"apple", "date"}
		result := sliceutil.FindDifferences(a, b)

		// Use ElementsMatch to ignore order
		assert.ElementsMatch(t, expected, result, "The unique elements should match regardless of order")
	})
}


func TestMergeSlices(t *testing.T) {
	// Test case 1: Merging slices of integers with ascending order
	t.Run("Merge Integers Ascending", func(t *testing.T) {
		a := []int{5, 1, 3}
		b := []int{4, 2, 6}

		expected := []int{1, 2, 3, 4, 5, 6}
		result := sliceutil.MergeSlices(a, b, sliceutil.ORDER_TYPE_ASC)
		assert.Equal(t, expected, result, "Merged slice should be sorted in ascending order")
	})

	// Test case 2: Merging slices of integers with descending order
	t.Run("Merge Integers Descending", func(t *testing.T) {
		a := []int{5, 1, 3}
		b := []int{4, 2, 6}

		expected := []int{6, 5, 4, 3, 2, 1}
		result := sliceutil.MergeSlices(a, b, sliceutil.ORDER_TYPE_DESC)
		assert.Equal(t, expected, result, "Merged slice should be sorted in descending order")
	})

	// Test case 3: Merging slices of strings with ascending order
	t.Run("Merge Strings Ascending", func(t *testing.T) {
		a := []string{"banana", "apple"}
		b := []string{"cherry", "date"}

		expected := []string{"apple", "banana", "cherry", "date"}
		result := sliceutil.MergeSlices(a, b, sliceutil.ORDER_TYPE_ASC)
		assert.Equal(t, expected, result, "Merged string slice should be sorted in ascending order")
	})

	// Test case 4: Merging slices of strings with descending order
	t.Run("Merge Strings Descending", func(t *testing.T) {
		a := []string{"banana", "apple"}
		b := []string{"cherry", "date"}

		expected := []string{"date", "cherry", "banana", "apple"}
		result := sliceutil.MergeSlices(a, b, sliceutil.ORDER_TYPE_DESC)
		assert.Equal(t, expected, result, "Merged string slice should be sorted in descending order")
	})

	// Test case 5: Merging empty slices
	t.Run("Merge Empty Slices", func(t *testing.T) {
		a := []int{}
		b := []int{}

		expected := []int{}
		result := sliceutil.MergeSlices(a, b, sliceutil.ORDER_TYPE_ASC)
		assert.Equal(t, expected, result, "Merging two empty slices should result in an empty slice")
	})

	// Test case 6: Merging a non-empty slice with an empty slice
	t.Run("Merge Non-Empty with Empty Slice", func(t *testing.T) {
		a := []int{3, 1, 4}
		b := []int{}

		expected := []int{1, 3, 4}
		result := sliceutil.MergeSlices(a, b, sliceutil.ORDER_TYPE_ASC)
		assert.Equal(t, expected, result, "Merging a non-empty slice with an empty slice should return the non-empty slice")
	})
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
