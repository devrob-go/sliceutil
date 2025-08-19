package sliceutil

import (
	"sort"
)

// FindDifferences returns unique values from both slices that are not in the other.
// This function uses a map-based approach to efficiently find symmetric differences.
//
// The function performs the following operations:
// 1. Creates a map to track element frequencies
// 2. Adds elements from slice A to the map
// 3. Processes slice B to find unique elements
// 4. Returns all elements that appear in only one of the slices
//
// Time complexity: O(n + m) where n and m are the lengths of the slices
// Space complexity: O(n + m) for the result slice
//
// Example:
//
//	a := []int{1, 2, 3, 4}
//	b := []int{3, 4, 5, 6}
//	result := FindDifferences(a, b) // returns []int{1, 2, 5, 6}
func FindDifferences[T comparable](a, b []T) []T {
	// Handle nil slices
	if a == nil && b == nil {
		return []T{}
	}
	if a == nil {
		return append([]T{}, b...)
	}
	if b == nil {
		return append([]T{}, a...)
	}

	// Create a map to track element frequencies
	m := make(map[T]int)

	// Add all elements from slice a to the map
	for _, v := range a {
		m[v]++
	}

	// Process slice b: if value exists, decrement; if not, add as unique
	for _, v := range b {
		if count, exists := m[v]; exists {
			if count == 1 {
				// Value exists in both slices, remove it
				delete(m, v)
			} else {
				// Value exists multiple times in slice a, decrement
				m[v]--
			}
		} else {
			// Value is unique to slice b
			m[v] = -1
		}
	}

	// Collect the remaining unique values into a result slice
	result := make([]T, 0, len(m))
	for k, v := range m {
		if v != 0 { // v != 0 means the element is unique to one slice
			result = append(result, k)
		}
	}

	return result
}

// FindDifferencesWithCount returns differences along with their frequency counts.
// This provides more detailed information about how many times each unique element appears.
func FindDifferencesWithCount[T comparable](a, b []T) map[T]int {
	// Handle nil slices
	if a == nil && b == nil {
		return make(map[T]int)
	}
	if a == nil {
		result := make(map[T]int)
		for _, v := range b {
			result[v] = -1 // Negative indicates unique to slice b
		}
		return result
	}
	if b == nil {
		result := make(map[T]int)
		for _, v := range a {
			result[v] = 1 // Positive indicates unique to slice a
		}
		return result
	}

	// Create a map to track element frequencies
	m := make(map[T]int)

	// Add all elements from slice a to the map
	for _, v := range a {
		m[v]++
	}

	// Process slice b: if value exists, decrement; if not, add as unique
	for _, v := range b {
		if count, exists := m[v]; exists {
			m[v] = count - 1
		} else {
			m[v] = -1 // Negative indicates unique to slice b
		}
	}

	// Filter out elements that appear in both slices with equal frequency
	result := make(map[T]int)
	for k, v := range m {
		if v != 0 {
			result[k] = v
		}
	}

	return result
}

// MaxInt returns the largest number in an int slice.
// The function returns an error if the slice is empty or nil.
//
// Time complexity: O(n) where n is the length of the slice
// Space complexity: O(1)
//
// Example:
//
//	slice := []int{1, 5, 3, 9, 2}
//	max, err := MaxInt(slice) // returns 9, nil
func MaxInt(a []int) (int, error) {
	if a == nil {
		return 0, ErrNilSlice
	}
	if len(a) == 0 {
		return 0, ErrEmptySlice
	}

	// Initialize max to the first element
	max := a[0]

	// Iterate through the slice and find the max value
	for _, v := range a {
		if v > max {
			// Update max if a larger element is found
			max = v
		}
	}
	return max, nil
}

// MinInt returns the smallest number in an int slice.
// The function returns an error if the slice is empty or nil.
//
// Time complexity: O(n) where n is the length of the slice
// Space complexity: O(1)
//
// Example:
//
//	slice := []int{1, 5, 3, 9, 2}
//	min, err := MinInt(slice) // returns 1, nil
func MinInt(a []int) (int, error) {
	if a == nil {
		return 0, ErrNilSlice
	}
	if len(a) == 0 {
		return 0, ErrEmptySlice
	}

	// Initialize min to the first element
	min := a[0]

	// Iterate through the slice and find the min value
	for _, v := range a {
		if v < min {
			// Update min if a smaller element is found
			min = v
		}
	}
	return min, nil
}

// MaxFloat64 returns the largest number in a float64 slice.
// The function returns an error if the slice is empty or nil.
func MaxFloat64(a []float64) (float64, error) {
	if a == nil {
		return 0, ErrNilSlice
	}
	if len(a) == 0 {
		return 0, ErrEmptySlice
	}

	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// MinFloat64 returns the smallest number in a float64 slice.
// The function returns an error if the slice is empty or nil.
func MinFloat64(a []float64) (float64, error) {
	if a == nil {
		return 0, ErrNilSlice
	}
	if len(a) == 0 {
		return 0, ErrEmptySlice
	}

	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min, nil
}

// SumInt calculates the sum of all integers in a slice.
// The function returns an error if the slice is nil.
func SumInt(a []int) (int, error) {
	if a == nil {
		return 0, ErrNilSlice
	}

	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum, nil
}

// SumFloat64 calculates the sum of all float64 values in a slice.
// The function returns an error if the slice is nil.
func SumFloat64(a []float64) (float64, error) {
	if a == nil {
		return 0, ErrNilSlice
	}

	sum := 0.0
	for _, v := range a {
		sum += v
	}
	return sum, nil
}

// AverageInt calculates the average of all integers in a slice.
// The function returns an error if the slice is empty or nil.
func AverageInt(a []int) (float64, error) {
	if a == nil {
		return 0, ErrNilSlice
	}
	if len(a) == 0 {
		return 0, ErrEmptySlice
	}

	sum, err := SumInt(a)
	if err != nil {
		return 0, err
	}

	return float64(sum) / float64(len(a)), nil
}

// AverageFloat64 calculates the average of all float64 values in a slice.
// The function returns an error if the slice is empty or nil.
func AverageFloat64(a []float64) (float64, error) {
	if a == nil {
		return 0, ErrNilSlice
	}
	if len(a) == 0 {
		return 0, ErrEmptySlice
	}

	sum, err := SumFloat64(a)
	if err != nil {
		return 0, err
	}

	return sum / float64(len(a)), nil
}

// CompareSum compares two int slices and determines which one has a greater sum.
// The function returns a Result type indicating the comparison outcome.
//
// Time complexity: O(n + m) where n and m are the lengths of the slices
// Space complexity: O(1)
//
// Example:
//
//	a := []int{1, 2, 3}
//	b := []int{4, 5, 6}
//	result := CompareSum(a, b) // returns ResultBGreater
func CompareSum(a, b []int) Result {
	// Calculate sum of elements in slice a
	sumA, errA := SumInt(a)
	if errA != nil {
		sumA = 0
	}

	// Calculate sum of elements in slice b
	sumB, errB := SumInt(b)
	if errB != nil {
		sumB = 0
	}

	// Compare the sums and return the result
	if sumA > sumB {
		return ResultAGreater
	} else if sumA < sumB {
		return ResultBGreater
	}
	return ResultEqual
}

// CompareSumWithDetails compares two int slices and provides detailed comparison results.
// This function is useful when you need more information about the comparison.
func CompareSumWithDetails(a, b []int) CompareResult {
	result := CompareResult{
		Details: make(map[string]interface{}),
	}

	// Calculate sums
	sumA, errA := SumInt(a)
	if errA != nil {
		result.Details["error_a"] = errA.Error()
		sumA = 0
	}

	sumB, errB := SumInt(b)
	if errB != nil {
		result.Details["error_b"] = errB.Error()
		sumB = 0
	}

	// Store sums in details
	result.Details["sum_a"] = sumA
	result.Details["sum_b"] = sumB
	result.Details["difference"] = sumA - sumB

	// Determine result
	if sumA > sumB {
		result.Equal = false
		result.Message = "Slice A has greater sum"
		result.Details["result"] = ResultAGreater
	} else if sumA < sumB {
		result.Equal = false
		result.Message = "Slice B has greater sum"
		result.Details["result"] = ResultBGreater
	} else {
		result.Equal = true
		result.Message = "Both slices have equal sums"
		result.Details["result"] = ResultEqual
	}

	return result
}

// GetSliceStats provides comprehensive statistical information about a slice.
// This function is useful for analyzing slice characteristics.
func GetSliceStats(a []int) (SliceStats, error) {
	if a == nil {
		return SliceStats{}, ErrNilSlice
	}

	stats := SliceStats{
		Length: len(a),
	}

	if len(a) == 0 {
		return stats, nil
	}

	// Calculate min, max, sum
	min, err := MinInt(a)
	if err != nil {
		return stats, err
	}
	stats.Min = min

	max, err := MaxInt(a)
	if err != nil {
		return stats, err
	}
	stats.Max = max

	sum, err := SumInt(a)
	if err != nil {
		return stats, err
	}
	stats.Sum = sum

	// Calculate average
	stats.Average = float64(sum) / float64(len(a))

	// Check for duplicates
	seen := make(map[int]bool)
	for _, v := range a {
		if seen[v] {
			stats.HasDuplicates = true
			break
		}
		seen[v] = true
	}

	return stats, nil
}

// IsSorted checks if a slice is sorted in ascending order.
// The function uses Go's sort.IsSorted for efficient checking.
func IsSorted[T sort.Interface](a T) bool {
	return sort.IsSorted(a)
}

// IsSortedInt checks if an int slice is sorted in ascending order.
func IsSortedInt(a []int) bool {
	if a == nil || len(a) <= 1 {
		return true
	}

	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

// IsSortedString checks if a string slice is sorted in ascending order.
func IsSortedString(a []string) bool {
	if a == nil || len(a) <= 1 {
		return true
	}

	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

// Reverse reverses the order of elements in a slice.
// The function modifies the original slice.
func Reverse[T any](a []T) {
	if a == nil || len(a) <= 1 {
		return
	}

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

// ReverseCopy creates a reversed copy of a slice without modifying the original.
func ReverseCopy[T any](a []T) []T {
	if a == nil {
		return nil
	}

	result := make([]T, len(a))
	copy(result, a)
	Reverse(result)
	return result
}

// RemoveDuplicates removes duplicate elements from a slice while preserving order.
// The function returns a new slice with duplicates removed.
func RemoveDuplicates[T comparable](a []T) []T {
	if a == nil {
		return nil
	}
	if len(a) <= 1 {
		return append([]T{}, a...)
	}

	seen := make(map[T]bool)
	result := make([]T, 0, len(a))

	for _, v := range a {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}

// Contains checks if a slice contains a specific element.
func Contains[T comparable](a []T, element T) bool {
	if a == nil {
		return false
	}

	for _, v := range a {
		if v == element {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of an element in a slice.
// Returns -1 if the element is not found.
func IndexOf[T comparable](a []T, element T) int {
	if a == nil {
		return -1
	}

	for i, v := range a {
		if v == element {
			return i
		}
	}
	return -1
}

// CountOccurrences counts how many times an element appears in a slice.
func CountOccurrences[T comparable](a []T, element T) int {
	if a == nil {
		return 0
	}

	count := 0
	for _, v := range a {
		if v == element {
			count++
		}
	}
	return count
}
