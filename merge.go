package sliceutil

import "sort"

const (
	ORDER_TYPE_ASC  = "ASC"
	ORDER_TYPE_DESC = "DESC"
)

// MergeSlices merges two slices of any type and sorts them based on the order (ascending or descending)
func MergeSlices(a, b interface{}, order string) interface{} {
	// Determine the type of the slices
	switch a := a.(type) {
	case []int:
		b := b.([]int)
		merged := append(a, b...)
		sort.Slice(merged, func(i, j int) bool {
			if order == ORDER_TYPE_ASC {
				return merged[i] < merged[j]
			}
			return merged[i] > merged[j]
		})
		return merged
	case []string:
		b := b.([]string)
		merged := append(a, b...)
		sort.Slice(merged, func(i, j int) bool {
			if order == ORDER_TYPE_ASC {
				return merged[i] < merged[j]
			}
			return merged[i] > merged[j]
		})
		return merged
	default:
		return nil
	}
}
