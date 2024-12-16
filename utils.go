package sliceutil

import "reflect"

// CompareSlices checks if two slices are equal in values and order (for int, string).
// It compares the length of both slices first, and if the lengths are different, they are not equal.
// Then, it iterates over the elements and compares them one by one. If any element doesn't match, it returns false.
// If all elements are the same, it returns true.
func CompareSlices[T comparable](a []T, b []T) bool {
	// Check if the lengths of the slices are different
	if len(a) != len(b) {
		return false
	}

	// Compare each element in the slices
	for i, v := range a {
		// If any element is different, the slices are not equal
		if v != b[i] {
			return false
		}
	}

	// If all elements are the same, return true
	return true
}

func CompareReflectionSlices(fieldA, fieldB reflect.Value) bool {
	// Ensure that the values are slices
	if fieldA.Kind() != reflect.Slice || fieldB.Kind() != reflect.Slice {
		return false
	}

	// We will check the type of slices and pass them accordingly to CompareSlices
	switch fieldA.Type().Elem().Kind() {
	case reflect.Int:
		// Type assertion for int slice
		a := fieldA.Interface().([]int)
		b := fieldB.Interface().([]int)
		return CompareSlices(a, b)
	case reflect.String:
		// Type assertion for string slice
		a := fieldA.Interface().([]string)
		b := fieldB.Interface().([]string)
		return CompareSlices(a, b)
	default:
		// If the type is unsupported, return false
		return false
	}
}

// FindDifferences returns unique values from both slices that are not in the other (for int and string).
// This function uses a map to track the elements of both slices, finding values that are not shared between them.
func FindDifferences[T comparable](a, b []T) []T {
	m := make(map[T]bool)

	// Add all elements from slice a to the map
	for _, v := range a {
		m[v] = true
	}

	// For slice b, if the value is already in the map, remove it, else add it for unique values
	for _, v := range b {
		if m[v] {
			// If value exists in both slices, remove it from the map
			delete(m, v)
		} else {
			// Otherwise, add the value as unique
			m[v] = true
		}
	}

	// Collect the remaining unique values into a result slice
	result := make([]T, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// MaxInt returns the largest number in an int slice.
// If the slice is empty, it panics. Otherwise, it iterates over the slice to find the maximum value.
func MaxInt(a []int) int {
	if len(a) == 0 {
		// We cannot find a max value in an empty slice
		panic("slice cannot be empty")
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
	return max
}

// MinInt returns the smallest number in an int slice.
// It functions similarly to MaxInt but looks for the minimum value in the slice.
func MinInt(a []int) int {
	if len(a) == 0 {
		// We cannot find a min value in an empty slice
		panic("slice cannot be empty")
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
	return min
}

// CompareSum compares two int slices and determines which one has a greater sum.
// It sums the elements of both slices and compares their totals.
func CompareSum(a, b []int) string {
	// Calculate sum of elements in slice a
	sumA, sumB := 0, 0
	for _, v := range a {
		sumA += v
	}

	// Calculate sum of elements in slice b
	for _, v := range b {
		sumB += v
	}

	// Compare the sums and return the result
	if sumA > sumB {
		return "a is greater"
	} else if sumA < sumB {
		return "b is greater"
	}
	return "both are equal"
}
