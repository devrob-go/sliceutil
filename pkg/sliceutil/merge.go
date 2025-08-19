package sliceutil

import (
	"sort"
)

// MergeSlices merges two slices of the same type and sorts them based on the specified order.
// The function supports int and string slices with ascending or descending sorting.
//
// The function performs the following operations:
// 1. Validates input parameters
// 2. Merges the slices
// 3. Sorts the merged result according to the specified order
// 4. Returns the sorted merged slice
//
// Time complexity: O((n + m) * log(n + m)) where n and m are the lengths of the slices
// Space complexity: O(n + m) for the result slice
//
// Example:
//
//	a := []int{5, 1, 3}
//	b := []int{4, 2, 6}
//	result := MergeSlices(a, b, OrderAsc) // returns []int{1, 2, 3, 4, 5, 6}
func MergeSlices(a, b interface{}, order OrderType) (interface{}, error) {
	// Validate order parameter
	if order != OrderAsc && order != OrderDesc {
		return nil, ErrUnsupportedType
	}

	// Determine the type of the slices and merge accordingly
	switch a := a.(type) {
	case []int:
		bSlice, ok := b.([]int)
		if !ok {
			return nil, ErrTypeMismatch
		}
		return mergeIntSlices(a, bSlice, order), nil
	case []string:
		bSlice, ok := b.([]string)
		if !ok {
			return nil, ErrTypeMismatch
		}
		return mergeStringSlices(a, bSlice, order), nil
	case []float64:
		bSlice, ok := b.([]float64)
		if !ok {
			return nil, ErrTypeMismatch
		}
		return mergeFloat64Slices(a, bSlice, order), nil
	default:
		return nil, ErrUnsupportedType
	}
}

// MergeSlicesGeneric is a generic version of MergeSlices that provides type safety
// for comparable types that can be sorted.
//
// This function requires the type parameter T to implement the sort.Interface,
// which means it must have a Less method for comparison.
func MergeSlicesGeneric[T any](a, b []T, order OrderType, less func(T, T) bool) []T {
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		result := make([]T, len(b))
		copy(result, b)
		sort.Slice(result, func(i, j int) bool {
			if order == OrderAsc {
				return less(result[i], result[j])
			}
			return less(result[j], result[i])
		})
		return result
	}
	if b == nil {
		result := make([]T, len(a))
		copy(result, a)
		sort.Slice(result, func(i, j int) bool {
			if order == OrderAsc {
				return less(result[i], result[j])
			}
			return less(result[j], result[i])
		})
		return result
	}

	// Merge slices
	merged := make([]T, 0, len(a)+len(b))
	merged = append(merged, a...)
	merged = append(merged, b...)

	// Sort merged slice
	sort.Slice(merged, func(i, j int) bool {
		if order == OrderAsc {
			return less(merged[i], merged[j])
		}
		return less(merged[j], merged[i])
	})

	return merged
}

// MergeSlicesInt merges two int slices with the specified sorting order.
// This function is a type-safe alternative to MergeSlices for int slices.
func MergeSlicesInt(a, b []int, order OrderType) []int {
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		result := make([]int, len(b))
		copy(result, b)
		sort.Ints(result)
		if order == OrderDesc {
			Reverse(result)
		}
		return result
	}
	if b == nil {
		result := make([]int, len(a))
		copy(result, a)
		sort.Ints(result)
		if order == OrderDesc {
			Reverse(result)
		}
		return result
	}

	// Merge slices
	merged := make([]int, 0, len(a)+len(b))
	merged = append(merged, a...)
	merged = append(merged, b...)

	// Sort merged slice
	if order == OrderAsc {
		sort.Ints(merged)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(merged)))
	}

	return merged
}

// MergeSlicesString merges two string slices with the specified sorting order.
// This function is a type-safe alternative to MergeSlices for string slices.
func MergeSlicesString(a, b []string, order OrderType) []string {
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		result := make([]string, len(b))
		copy(result, b)
		sort.Strings(result)
		if order == OrderDesc {
			Reverse(result)
		}
		return result
	}
	if b == nil {
		result := make([]string, len(a))
		copy(result, a)
		sort.Strings(result)
		if order == OrderDesc {
			Reverse(result)
		}
		return result
	}

	// Merge slices
	merged := make([]string, 0, len(a)+len(b))
	merged = append(merged, a...)
	merged = append(merged, b...)

	// Sort merged slice
	if order == OrderAsc {
		sort.Strings(merged)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(merged)))
	}

	return merged
}

// MergeSlicesFloat64 merges two float64 slices with the specified sorting order.
// This function is a type-safe alternative to MergeSlices for float64 slices.
func MergeSlicesFloat64(a, b []float64, order OrderType) []float64 {
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		result := make([]float64, len(b))
		copy(result, b)
		sort.Float64s(result)
		if order == OrderDesc {
			Reverse(result)
		}
		return result
	}
	if b == nil {
		result := make([]float64, len(a))
		copy(result, a)
		sort.Float64s(result)
		if order == OrderDesc {
			Reverse(result)
		}
		return result
	}

	// Merge slices
	merged := make([]float64, 0, len(a)+len(b))
	merged = append(merged, a...)
	merged = append(merged, b...)

	// Sort merged slice
	if order == OrderAsc {
		sort.Float64s(merged)
	} else {
		sort.Sort(sort.Reverse(sort.Float64Slice(merged)))
	}

	return merged
}

// MergeMultipleSlices merges multiple slices of the same type and sorts them.
// This function is useful when you need to merge more than two slices.
func MergeMultipleSlices[T any](slices [][]T, order OrderType, less func(T, T) bool) []T {
	if len(slices) == 0 {
		return nil
	}
	if len(slices) == 1 {
		result := make([]T, len(slices[0]))
		copy(result, slices[0])
		// Sort the single slice according to the order
		sort.Slice(result, func(i, j int) bool {
			if order == OrderAsc {
				return less(result[i], result[j])
			}
			return less(result[j], result[i])
		})
		return result
	}

	// Calculate total capacity
	totalCap := 0
	for _, slice := range slices {
		if slice != nil {
			totalCap += len(slice)
		}
	}

	// Merge all slices
	merged := make([]T, 0, totalCap)
	for _, slice := range slices {
		if slice != nil {
			merged = append(merged, slice...)
		}
	}

	// Sort merged slice
	sort.Slice(merged, func(i, j int) bool {
		if order == OrderAsc {
			return less(merged[i], merged[j])
		}
		return less(merged[j], merged[i])
	})

	return merged
}

// MergeSlicesWithDeduplication merges two slices and removes duplicates.
// This function is useful when you want to merge slices while ensuring uniqueness.
func MergeSlicesWithDeduplication[T comparable](a, b []T, order OrderType, less func(T, T) bool) []T {
	if a == nil && b == nil {
		return nil
	}

	// Merge slices
	var merged []T
	if a == nil {
		merged = append([]T{}, b...)
	} else if b == nil {
		merged = append([]T{}, a...)
	} else {
		merged = make([]T, 0, len(a)+len(b))
		merged = append(merged, a...)
		merged = append(merged, b...)
	}

	// Remove duplicates
	merged = RemoveDuplicates(merged)

	// Sort merged slice
	sort.Slice(merged, func(i, j int) bool {
		if order == OrderAsc {
			return less(merged[i], merged[j])
		}
		return less(merged[j], merged[i])
	})

	return merged
}

// Helper functions for specific types

// mergeIntSlices is a helper function that merges int slices
func mergeIntSlices(a, b []int, order OrderType) []int {
	return MergeSlicesInt(a, b, order)
}

// mergeStringSlices is a helper function that merges string slices
func mergeStringSlices(a, b []string, order OrderType) []string {
	return MergeSlicesString(a, b, order)
}

// mergeFloat64Slices is a helper function that merges float64 slices
func mergeFloat64Slices(a, b []float64, order OrderType) []float64 {
	return MergeSlicesFloat64(a, b, order)
}

// MergeSlicesWithCustomSort merges two slices using a custom sorting function.
// This function provides maximum flexibility for custom sorting logic.
func MergeSlicesWithCustomSort[T any](a, b []T, sortFunc func([]T)) []T {
	if a == nil && b == nil {
		return nil
	}

	// Merge slices
	var merged []T
	if a == nil {
		merged = append([]T{}, b...)
	} else if b == nil {
		merged = append([]T{}, a...)
	} else {
		merged = make([]T, 0, len(a)+len(b))
		merged = append(merged, a...)
		merged = append(merged, b...)
	}

	// Apply custom sorting
	sortFunc(merged)

	return merged
}

// MergeSlicesWithStableSort merges two slices using a stable sort algorithm.
// Stable sort preserves the relative order of equal elements.
func MergeSlicesWithStableSort[T any](a, b []T, order OrderType, less func(T, T) bool) []T {
	if a == nil && b == nil {
		return nil
	}

	// Merge slices
	var merged []T
	if a == nil {
		merged = append([]T{}, b...)
	} else if b == nil {
		merged = append([]T{}, a...)
	} else {
		merged = make([]T, 0, len(a)+len(b))
		merged = append(merged, a...)
		merged = append(merged, b...)
	}

	// Apply stable sort
	sort.SliceStable(merged, func(i, j int) bool {
		if order == OrderAsc {
			return less(merged[i], merged[j])
		}
		return less(merged[j], merged[i])
	})

	return merged
}
