package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMergeSlices tests the MergeSlices function
func TestMergeSlices(t *testing.T) {
	t.Run("Merge Int Slices Ascending", func(t *testing.T) {
		a := []int{5, 1, 3}
		b := []int{4, 2, 6}
		expected := []int{1, 2, 3, 4, 5, 6}

		result, err := MergeSlices(a, b, OrderAsc)
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Merge Int Slices Descending", func(t *testing.T) {
		a := []int{5, 1, 3}
		b := []int{4, 2, 6}
		expected := []int{6, 5, 4, 3, 2, 1}

		result, err := MergeSlices(a, b, OrderDesc)
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Merge String Slices Ascending", func(t *testing.T) {
		a := []string{"banana", "apple"}
		b := []string{"cherry", "date"}
		expected := []string{"apple", "banana", "cherry", "date"}

		result, err := MergeSlices(a, b, OrderAsc)
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Merge String Slices Descending", func(t *testing.T) {
		a := []string{"banana", "apple"}
		b := []string{"cherry", "date"}
		expected := []string{"date", "cherry", "banana", "apple"}

		result, err := MergeSlices(a, b, OrderDesc)
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Merge Float64 Slices", func(t *testing.T) {
		a := []float64{3.3, 1.1}
		b := []float64{4.4, 2.2}
		expected := []float64{1.1, 2.2, 3.3, 4.4}

		result, err := MergeSlices(a, b, OrderAsc)
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Invalid Order Type", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}

		_, err := MergeSlices(a, b, "INVALID")
		assert.ErrorIs(t, err, ErrUnsupportedType)
	})

	t.Run("Type Mismatch", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []string{"a", "b", "c"}

		_, err := MergeSlices(a, b, OrderAsc)
		assert.ErrorIs(t, err, ErrTypeMismatch)
	})

	t.Run("Unsupported Type", func(t *testing.T) {
		a := []bool{true, false}
		b := []bool{false, true}

		_, err := MergeSlices(a, b, OrderAsc)
		assert.ErrorIs(t, err, ErrUnsupportedType)
	})
}

// TestMergeSlicesGeneric tests the generic merge function
func TestMergeSlicesGeneric(t *testing.T) {
	t.Run("Merge with Custom Less Function", func(t *testing.T) {
		a := []int{5, 1, 3}
		b := []int{4, 2, 6}
		less := func(a, b int) bool { return a < b }

		result := MergeSlicesGeneric(a, b, OrderAsc, less)
		expected := []int{1, 2, 3, 4, 5, 6}
		assert.Equal(t, expected, result)
	})

	t.Run("Merge with Descending Order", func(t *testing.T) {
		a := []int{5, 1, 3}
		b := []int{4, 2, 6}
		less := func(a, b int) bool { return a < b }

		result := MergeSlicesGeneric(a, b, OrderDesc, less)
		expected := []int{6, 5, 4, 3, 2, 1}
		assert.Equal(t, expected, result)
	})

	t.Run("Nil Slices", func(t *testing.T) {
		less := func(a, b int) bool { return a < b }

		assert.Nil(t, MergeSlicesGeneric[int](nil, nil, OrderAsc, less))
		assert.Equal(t, []int{1, 2, 3}, MergeSlicesGeneric(nil, []int{1, 2, 3}, OrderAsc, less))
		assert.Equal(t, []int{1, 2, 3}, MergeSlicesGeneric([]int{1, 2, 3}, nil, OrderAsc, less))
	})
}

// TestMergeSlicesInt tests the type-safe int merge function
func TestMergeSlicesInt(t *testing.T) {
	t.Run("Merge Int Slices Ascending", func(t *testing.T) {
		a := []int{5, 1, 3}
		b := []int{4, 2, 6}
		expected := []int{1, 2, 3, 4, 5, 6}

		result := MergeSlicesInt(a, b, OrderAsc)
		assert.Equal(t, expected, result)
	})

	t.Run("Merge Int Slices Descending", func(t *testing.T) {
		a := []int{5, 1, 3}
		b := []int{4, 2, 6}
		expected := []int{6, 5, 4, 3, 2, 1}

		result := MergeSlicesInt(a, b, OrderDesc)
		assert.Equal(t, expected, result)
	})

	t.Run("Nil Slices", func(t *testing.T) {
		assert.Nil(t, MergeSlicesInt(nil, nil, OrderAsc))
		assert.Equal(t, []int{1, 2, 3}, MergeSlicesInt(nil, []int{1, 2, 3}, OrderAsc))
		assert.Equal(t, []int{1, 2, 3}, MergeSlicesInt([]int{1, 2, 3}, nil, OrderAsc))
	})
}

// TestMergeSlicesString tests the type-safe string merge function
func TestMergeSlicesString(t *testing.T) {
	t.Run("Merge String Slices Ascending", func(t *testing.T) {
		a := []string{"banana", "apple"}
		b := []string{"cherry", "date"}
		expected := []string{"apple", "banana", "cherry", "date"}

		result := MergeSlicesString(a, b, OrderAsc)
		assert.Equal(t, expected, result)
	})

	t.Run("Merge String Slices Descending", func(t *testing.T) {
		a := []string{"banana", "apple"}
		b := []string{"cherry", "date"}
		expected := []string{"date", "cherry", "banana", "apple"}

		result := MergeSlicesString(a, b, OrderDesc)
		assert.Equal(t, expected, result)
	})

	t.Run("Nil Slices", func(t *testing.T) {
		assert.Nil(t, MergeSlicesString(nil, nil, OrderAsc))
		assert.Equal(t, []string{"a", "b"}, MergeSlicesString(nil, []string{"a", "b"}, OrderAsc))
		assert.Equal(t, []string{"a", "b"}, MergeSlicesString([]string{"a", "b"}, nil, OrderAsc))
	})
}

// TestMergeSlicesFloat64 tests the type-safe float64 merge function
func TestMergeSlicesFloat64(t *testing.T) {
	t.Run("Merge Float64 Slices Ascending", func(t *testing.T) {
		a := []float64{3.3, 1.1}
		b := []float64{4.4, 2.2}
		expected := []float64{1.1, 2.2, 3.3, 4.4}

		result := MergeSlicesFloat64(a, b, OrderAsc)
		assert.Equal(t, expected, result)
	})

	t.Run("Merge Float64 Slices Descending", func(t *testing.T) {
		a := []float64{3.3, 1.1}
		b := []float64{4.4, 2.2}
		expected := []float64{4.4, 3.3, 2.2, 1.1}

		result := MergeSlicesFloat64(a, b, OrderDesc)
		assert.Equal(t, expected, result)
	})

	t.Run("Nil Slices", func(t *testing.T) {
		assert.Nil(t, MergeSlicesFloat64(nil, nil, OrderAsc))
		assert.Equal(t, []float64{1.1, 2.2}, MergeSlicesFloat64(nil, []float64{1.1, 2.2}, OrderAsc))
		assert.Equal(t, []float64{1.1, 2.2}, MergeSlicesFloat64([]float64{1.1, 2.2}, nil, OrderAsc))
	})
}

// TestMergeMultipleSlices tests merging multiple slices
func TestMergeMultipleSlices(t *testing.T) {
	t.Run("Merge Three Slices", func(t *testing.T) {
		slices := [][]int{
			{3, 1},
			{4, 2},
			{6, 5},
		}
		less := func(a, b int) bool { return a < b }

		result := MergeMultipleSlices(slices, OrderAsc, less)
		expected := []int{1, 2, 3, 4, 5, 6}
		assert.Equal(t, expected, result)
	})

	t.Run("Merge with Nil Slices", func(t *testing.T) {
		slices := [][]int{
			{3, 1},
			nil,
			{4, 2},
		}
		less := func(a, b int) bool { return a < b }

		result := MergeMultipleSlices(slices, OrderAsc, less)
		expected := []int{1, 2, 3, 4}
		assert.Equal(t, expected, result)
	})

	t.Run("Empty Slices Array", func(t *testing.T) {
		var slices [][]int
		less := func(a, b int) bool { return a < b }

		result := MergeMultipleSlices(slices, OrderAsc, less)
		assert.Nil(t, result)
	})

	t.Run("Single Slice", func(t *testing.T) {
		slices := [][]int{{3, 1, 2}}
		less := func(a, b int) bool { return a < b }

		result := MergeMultipleSlices(slices, OrderAsc, less)
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})
}

// TestMergeSlicesWithDeduplication tests merging with duplicate removal
func TestMergeSlicesWithDeduplication(t *testing.T) {
	t.Run("Merge with Duplicates", func(t *testing.T) {
		a := []int{1, 2, 2, 3}
		b := []int{3, 4, 4, 5}
		less := func(a, b int) bool { return a < b }

		result := MergeSlicesWithDeduplication(a, b, OrderAsc, less)
		expected := []int{1, 2, 3, 4, 5}
		assert.Equal(t, expected, result)
	})

	t.Run("No Duplicates", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}
		less := func(a, b int) bool { return a < b }

		result := MergeSlicesWithDeduplication(a, b, OrderAsc, less)
		expected := []int{1, 2, 3, 4, 5, 6}
		assert.Equal(t, expected, result)
	})
}

// TestMergeSlicesWithCustomSort tests custom sorting
func TestMergeSlicesWithCustomSort(t *testing.T) {
	t.Run("Custom Sort Function", func(t *testing.T) {
		a := []int{5, 1, 3}
		b := []int{4, 2, 6}

		sortFunc := func(slice []int) {
			// Custom reverse sort
			for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}

		result := MergeSlicesWithCustomSort(a, b, sortFunc)
		expected := []int{6, 2, 4, 3, 1, 5}
		assert.Equal(t, expected, result)
	})
}

// TestMergeSlicesWithStableSort tests stable sorting
func TestMergeSlicesWithStableSort(t *testing.T) {
	t.Run("Stable Sort", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		less := func(a, b int) bool { return a < b }

		result := MergeSlicesWithStableSort(a, b, OrderAsc, less)
		expected := []int{1, 1, 2, 2, 3, 3}
		assert.Equal(t, expected, result)
	})
}

// TestEdgeCases tests various edge cases
func TestEdgeCases(t *testing.T) {
	t.Run("Empty Slices", func(t *testing.T) {
		a := []int{}
		b := []int{}

		result, err := MergeSlices(a, b, OrderAsc)
		require.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("One Empty Slice", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{}

		result, err := MergeSlices(a, b, OrderAsc)
		require.NoError(t, err)
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("Large Slices", func(t *testing.T) {
		a := make([]int, 1000)
		b := make([]int, 1000)

		for i := range a {
			a[i] = 1000 - i
			b[i] = i + 1
		}

		result, err := MergeSlices(a, b, OrderAsc)
		require.NoError(t, err)
		resultSlice := result.([]int)
		assert.Equal(t, 2000, len(resultSlice))
		assert.True(t, IsSortedInt(resultSlice))
	})
}
