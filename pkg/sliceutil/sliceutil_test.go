package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCompareSlices tests the CompareSlices function with various scenarios
func TestCompareSlices(t *testing.T) {
	t.Run("Identical Int Slices", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{1, 2, 3, 4, 5}
		assert.True(t, CompareSlices(a, b))
	})

	t.Run("Different Order Int Slices", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{5, 4, 3, 2, 1}
		assert.False(t, CompareSlices(a, b))
	})

	t.Run("Different Length Int Slices", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3, 4}
		assert.False(t, CompareSlices(a, b))
	})

	t.Run("Identical String Slices", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		b := []string{"apple", "banana", "cherry"}
		assert.True(t, CompareSlices(a, b))
	})

	t.Run("Different String Slices", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		b := []string{"apple", "orange", "cherry"}
		assert.False(t, CompareSlices(a, b))
	})

	t.Run("Nil Slices", func(t *testing.T) {
		assert.True(t, CompareSlices[int](nil, nil))
		assert.False(t, CompareSlices([]int{1, 2, 3}, nil))
		assert.False(t, CompareSlices[int](nil, []int{1, 2, 3}))
	})

	t.Run("Empty Slices", func(t *testing.T) {
		a := []int{}
		b := []int{}
		assert.True(t, CompareSlices(a, b))
	})
}

// TestCompareSlicesWithResult tests the CompareSlicesWithResult function
func TestCompareSlicesWithResult(t *testing.T) {
	t.Run("Equal Slices", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		result := CompareSlicesWithResult(a, b)

		assert.True(t, result.Equal)
		assert.Equal(t, "Slices are equal", result.Message)
		assert.Empty(t, result.Details)
	})

	t.Run("Different Length Slices", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2}
		result := CompareSlicesWithResult(a, b)

		assert.False(t, result.Equal)
		assert.Equal(t, "Slices have different lengths", result.Message)
		assert.Equal(t, 3, result.Details["length_a"])
		assert.Equal(t, 2, result.Details["length_b"])
	})

	t.Run("Different Values Slices", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 5, 3}
		result := CompareSlicesWithResult(a, b)

		assert.False(t, result.Equal)
		assert.Equal(t, "Slices differ at specific indices", result.Message)
		assert.Equal(t, 1, result.Details["difference_count"])
		assert.Equal(t, []int{1}, result.Details["differences"])
	})

	t.Run("Nil Slices", func(t *testing.T) {
		result := CompareSlicesWithResult[int](nil, []int{1, 2, 3})

		assert.False(t, result.Equal)
		assert.Equal(t, "One slice is nil while the other is not", result.Message)
		assert.True(t, result.Details["a_nil"].(bool))
		assert.False(t, result.Details["b_nil"].(bool))
	})
}

// TestFindDifferences tests the FindDifferences function
func TestFindDifferences(t *testing.T) {
	t.Run("Int Slices with Differences", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		b := []int{3, 4, 5, 6}
		expected := []int{1, 2, 5, 6}
		result := FindDifferences(a, b)

		// Use ElementsMatch to ignore order
		assert.ElementsMatch(t, expected, result)
	})

	t.Run("String Slices with Differences", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		b := []string{"banana", "date", "cherry"}
		expected := []string{"apple", "date"}
		result := FindDifferences(a, b)

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("Nil Slices", func(t *testing.T) {
		assert.Empty(t, FindDifferences[int](nil, nil))
		assert.Equal(t, []int{1, 2, 3}, FindDifferences[int](nil, []int{1, 2, 3}))
		assert.Equal(t, []int{1, 2, 3}, FindDifferences([]int{1, 2, 3}, nil))
	})

	t.Run("Empty Slices", func(t *testing.T) {
		a := []int{}
		b := []int{}
		assert.Empty(t, FindDifferences(a, b))
	})
}

// TestFindDifferencesWithCount tests the FindDifferencesWithCount function
func TestFindDifferencesWithCount(t *testing.T) {
	t.Run("Int Slices with Counts", func(t *testing.T) {
		a := []int{1, 2, 2, 3}
		b := []int{2, 3, 4, 4}
		result := FindDifferencesWithCount(a, b)

		expected := map[int]int{
			1: 1,  // Unique to slice a
			2: 1,  // Appears twice in a, once in b (difference: 1)
			4: -2, // Unique to slice b (appears twice)
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Nil Slices with Counts", func(t *testing.T) {
		result := FindDifferencesWithCount[int](nil, []int{1, 2, 3})
		expected := map[int]int{
			1: -1,
			2: -1,
			3: -1,
		}
		assert.Equal(t, expected, result)
	})
}

// TestMaxMinInt tests the MaxInt and MinInt functions
func TestMaxMinInt(t *testing.T) {
	t.Run("MaxInt Success", func(t *testing.T) {
		slice := []int{1, 5, 3, 9, 2}
		max, err := MaxInt(slice)

		require.NoError(t, err)
		assert.Equal(t, 9, max)
	})

	t.Run("MaxInt Empty Slice", func(t *testing.T) {
		_, err := MaxInt([]int{})
		assert.ErrorIs(t, err, ErrEmptySlice)
	})

	t.Run("MaxInt Nil Slice", func(t *testing.T) {
		_, err := MaxInt(nil)
		assert.ErrorIs(t, err, ErrNilSlice)
	})

	t.Run("MinInt Success", func(t *testing.T) {
		slice := []int{1, 5, 3, 9, 2}
		min, err := MinInt(slice)

		require.NoError(t, err)
		assert.Equal(t, 1, min)
	})

	t.Run("MinInt Empty Slice", func(t *testing.T) {
		_, err := MinInt([]int{})
		assert.ErrorIs(t, err, ErrEmptySlice)
	})

	t.Run("MinInt Nil Slice", func(t *testing.T) {
		_, err := MinInt(nil)
		assert.ErrorIs(t, err, ErrNilSlice)
	})
}

// TestFloat64Functions tests the float64 utility functions
func TestFloat64Functions(t *testing.T) {
	t.Run("MaxFloat64", func(t *testing.T) {
		slice := []float64{1.1, 5.5, 3.3, 9.9, 2.2}
		max, err := MaxFloat64(slice)

		require.NoError(t, err)
		assert.Equal(t, 9.9, max)
	})

	t.Run("MinFloat64", func(t *testing.T) {
		slice := []float64{1.1, 5.5, 3.3, 9.9, 2.2}
		min, err := MinFloat64(slice)

		require.NoError(t, err)
		assert.Equal(t, 1.1, min)
	})

	t.Run("SumFloat64", func(t *testing.T) {
		slice := []float64{1.0, 2.0, 3.0}
		sum, err := SumFloat64(slice)

		require.NoError(t, err)
		assert.Equal(t, 6.0, sum)
	})

	t.Run("AverageFloat64", func(t *testing.T) {
		slice := []float64{1.0, 2.0, 3.0, 4.0}
		avg, err := AverageFloat64(slice)

		require.NoError(t, err)
		assert.Equal(t, 2.5, avg)
	})
}

// TestSumAndAverage tests the sum and average functions
func TestSumAndAverage(t *testing.T) {
	t.Run("SumInt", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		sum, err := SumInt(slice)

		require.NoError(t, err)
		assert.Equal(t, 15, sum)
	})

	t.Run("SumInt Nil Slice", func(t *testing.T) {
		_, err := SumInt(nil)
		assert.ErrorIs(t, err, ErrNilSlice)
	})

	t.Run("AverageInt", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		avg, err := AverageInt(slice)

		require.NoError(t, err)
		assert.Equal(t, 3.0, avg)
	})

	t.Run("AverageInt Empty Slice", func(t *testing.T) {
		_, err := AverageInt([]int{})
		assert.ErrorIs(t, err, ErrEmptySlice)
	})
}

// TestCompareSum tests the CompareSum function
func TestCompareSum(t *testing.T) {
	t.Run("Slice A Greater", func(t *testing.T) {
		a := []int{10, 20}
		b := []int{5, 5}
		result := CompareSum(a, b)
		assert.Equal(t, ResultAGreater, result)
	})

	t.Run("Slice B Greater", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}
		result := CompareSum(a, b)
		assert.Equal(t, ResultBGreater, result)
	})

	t.Run("Equal Sums", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{2, 1, 3}
		result := CompareSum(a, b)
		assert.Equal(t, ResultEqual, result)
	})

	t.Run("Nil Slices", func(t *testing.T) {
		result := CompareSum(nil, []int{1, 2, 3})
		assert.Equal(t, ResultBGreater, result)
	})
}

// TestCompareSumWithDetails tests the CompareSumWithDetails function
func TestCompareSumWithDetails(t *testing.T) {
	t.Run("Detailed Comparison", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}
		result := CompareSumWithDetails(a, b)

		assert.False(t, result.Equal)
		assert.Equal(t, "Slice B has greater sum", result.Message)
		assert.Equal(t, 6, result.Details["sum_a"])
		assert.Equal(t, 15, result.Details["sum_b"])
		assert.Equal(t, -9, result.Details["difference"])
		assert.Equal(t, ResultBGreater, result.Details["result"])
	})
}

// TestGetSliceStats tests the GetSliceStats function
func TestGetSliceStats(t *testing.T) {
	t.Run("Slice Statistics", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		stats, err := GetSliceStats(slice)

		require.NoError(t, err)
		assert.Equal(t, 5, stats.Length)
		assert.Equal(t, 1, stats.Min)
		assert.Equal(t, 5, stats.Max)
		assert.Equal(t, 15, stats.Sum)
		assert.Equal(t, 3.0, stats.Average)
		assert.False(t, stats.HasDuplicates)
	})

	t.Run("Slice with Duplicates", func(t *testing.T) {
		slice := []int{1, 2, 2, 3}
		stats, err := GetSliceStats(slice)

		require.NoError(t, err)
		assert.True(t, stats.HasDuplicates)
	})

	t.Run("Empty Slice", func(t *testing.T) {
		stats, err := GetSliceStats([]int{})

		require.NoError(t, err)
		assert.Equal(t, 0, stats.Length)
		assert.False(t, stats.HasDuplicates)
	})

	t.Run("Nil Slice", func(t *testing.T) {
		_, err := GetSliceStats(nil)
		assert.ErrorIs(t, err, ErrNilSlice)
	})
}

// TestSortingFunctions tests the sorting utility functions
func TestSortingFunctions(t *testing.T) {
	t.Run("IsSortedInt", func(t *testing.T) {
		assert.True(t, IsSortedInt([]int{1, 2, 3, 4, 5}))
		assert.False(t, IsSortedInt([]int{5, 4, 3, 2, 1}))
		assert.True(t, IsSortedInt([]int{}))
		assert.True(t, IsSortedInt(nil))
	})

	t.Run("IsSortedString", func(t *testing.T) {
		assert.True(t, IsSortedString([]string{"a", "b", "c"}))
		assert.False(t, IsSortedString([]string{"c", "b", "a"}))
		assert.True(t, IsSortedString([]string{}))
		assert.True(t, IsSortedString(nil))
	})
}

// TestReverseFunctions tests the reverse utility functions
func TestReverseFunctions(t *testing.T) {
	t.Run("Reverse", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		Reverse(slice)
		assert.Equal(t, []int{5, 4, 3, 2, 1}, slice)
	})

	t.Run("ReverseCopy", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		reversed := ReverseCopy(original)

		assert.Equal(t, []int{5, 4, 3, 2, 1}, reversed)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, original) // Original unchanged
	})

	t.Run("Reverse Nil and Empty", func(t *testing.T) {
		Reverse[int](nil) // Should not panic
		Reverse([]int{})  // Should not panic
		Reverse([]int{1}) // Should not panic
	})
}

// TestRemoveDuplicates tests the RemoveDuplicates function
func TestRemoveDuplicates(t *testing.T) {
	t.Run("Remove Duplicates", func(t *testing.T) {
		slice := []int{1, 2, 2, 3, 3, 4}
		result := RemoveDuplicates(slice)

		assert.Equal(t, []int{1, 2, 3, 4}, result)
		assert.Equal(t, []int{1, 2, 2, 3, 3, 4}, slice) // Original unchanged
	})

	t.Run("No Duplicates", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		result := RemoveDuplicates(slice)

		assert.Equal(t, slice, result)
	})

	t.Run("Nil and Empty", func(t *testing.T) {
		assert.Nil(t, RemoveDuplicates[int](nil))
		assert.Empty(t, RemoveDuplicates([]int{}))
		assert.Equal(t, []int{1}, RemoveDuplicates([]int{1}))
	})
}

// TestSearchFunctions tests the search utility functions
func TestSearchFunctions(t *testing.T) {
	t.Run("Contains", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		assert.True(t, Contains(slice, 3))
		assert.False(t, Contains(slice, 6))
		assert.False(t, Contains[int](nil, 1))
	})

	t.Run("IndexOf", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		assert.Equal(t, 2, IndexOf(slice, 3))
		assert.Equal(t, -1, IndexOf(slice, 6))
		assert.Equal(t, -1, IndexOf[int](nil, 1))
	})

	t.Run("CountOccurrences", func(t *testing.T) {
		slice := []int{1, 2, 2, 3, 2, 4}
		assert.Equal(t, 3, CountOccurrences(slice, 2))
		assert.Equal(t, 1, CountOccurrences(slice, 1))
		assert.Equal(t, 0, CountOccurrences(slice, 6))
		assert.Equal(t, 0, CountOccurrences[int](nil, 1))
	})
}

// TestErrorConstants tests that error constants are properly defined
func TestErrorConstants(t *testing.T) {
	assert.NotNil(t, ErrEmptySlice)
	assert.NotNil(t, ErrNilSlice)
	assert.NotNil(t, ErrTypeMismatch)
	assert.NotNil(t, ErrUnsupportedType)

	assert.Equal(t, "slice cannot be empty", ErrEmptySlice.Error())
	assert.Equal(t, "slice cannot be nil", ErrNilSlice.Error())
	assert.Equal(t, "slice types do not match", ErrTypeMismatch.Error())
	assert.Equal(t, "unsupported slice type", ErrUnsupportedType.Error())
}

// TestOrderTypeConstants tests that order type constants are properly defined
func TestOrderTypeConstants(t *testing.T) {
	assert.Equal(t, OrderType("ASC"), OrderAsc)
	assert.Equal(t, OrderType("DESC"), OrderDesc)
}

// TestResultConstants tests that result constants are properly defined
func TestResultConstants(t *testing.T) {
	assert.Equal(t, Result("a is greater"), ResultAGreater)
	assert.Equal(t, Result("b is greater"), ResultBGreater)
	assert.Equal(t, Result("both are equal"), ResultEqual)
}
