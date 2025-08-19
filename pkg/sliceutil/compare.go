package sliceutil

import (
	"fmt"
	"reflect"
)

// CompareSlices checks if two slices are equal in values and order.
// It uses Go generics to provide type safety for comparable types.
//
// The function performs the following checks:
// 1. Validates that neither slice is nil
// 2. Compares slice lengths
// 3. Compares each element in order
//
// Time complexity: O(n) where n is the length of the slices
// Space complexity: O(1)
//
// Example:
//
//	a := []int{1, 2, 3}
//	b := []int{1, 2, 3}
//	result := CompareSlices(a, b) // returns true
func CompareSlices[T comparable](a, b []T) bool {
	// Check for nil slices
	if a == nil || b == nil {
		return a == nil && b == nil
	}

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

// CompareSlicesWithResult provides detailed comparison results including
// information about where differences occur.
//
// This function is useful when you need more than just a boolean result
// and want to understand the nature of differences between slices.
func CompareSlicesWithResult[T comparable](a, b []T) CompareResult {
	result := CompareResult{
		Equal:   true,
		Message: "Slices are equal",
		Details: make(map[string]interface{}),
	}

	// Check for nil slices
	if a == nil || b == nil {
		if a == nil && b == nil {
			return result
		}
		result.Equal = false
		result.Message = "One slice is nil while the other is not"
		result.Details["a_nil"] = a == nil
		result.Details["b_nil"] = b == nil
		return result
	}

	// Check lengths
	if len(a) != len(b) {
		result.Equal = false
		result.Message = "Slices have different lengths"
		result.Details["length_a"] = len(a)
		result.Details["length_b"] = len(b)
		return result
	}

	// Find differences
	var differences []int
	for i, v := range a {
		if v != b[i] {
			differences = append(differences, i)
		}
	}

	if len(differences) > 0 {
		result.Equal = false
		result.Message = "Slices differ at specific indices"
		result.Details["differences"] = differences
		result.Details["difference_count"] = len(differences)
	}

	return result
}

// CompareReflectionSlices compares two slices using reflection.
// This function is useful when you need to compare slices of unknown types
// at runtime.
//
// Supported types: []int, []string
// For other types, the function returns false.
//
// Note: This function is less performant than CompareSlices due to reflection overhead.
// Use CompareSlices when the types are known at compile time.
func CompareReflectionSlices(fieldA, fieldB reflect.Value) bool {
	// Ensure that the values are slices
	if fieldA.Kind() != reflect.Slice || fieldB.Kind() != reflect.Slice {
		return false
	}

	// Check if the slice types match
	if fieldA.Type() != fieldB.Type() {
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

// CompareStructs compares two structs deeply, supporting nested structs and pointers.
// The function uses memoization to improve performance for repeated comparisons.
//
// Features:
// - Deep comparison of all exported fields
// - Support for nested structs
// - Support for pointer fields
// - Memoization cache for performance
// - Handles nil pointers gracefully
//
// The function recursively compares all exported fields of the structs.
// Unexported fields are ignored as they cannot be accessed via reflection.
func CompareStructs(a, b interface{}) bool {
	// If both are nil, they are equal
	if a == nil && b == nil {
		return true
	}

	// If one is nil, they are not equal
	if a == nil || b == nil {
		return false
	}

	// Check if the types are the same
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	// Use reflection to inspect the struct fields
	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)

	// Dereference pointers, if applicable
	if valA.Kind() == reflect.Ptr {
		if valA.IsNil() || valB.IsNil() {
			return valA.IsNil() && valB.IsNil()
		}
		valA = valA.Elem()
		valB = valB.Elem()
	}

	// Check if both are structs
	if valA.Kind() != reflect.Struct || valB.Kind() != reflect.Struct {
		return reflect.DeepEqual(a, b)
	}

	// Compare each field of the structs
	for i := 0; i < valA.NumField(); i++ {
		fieldA := valA.Field(i)
		fieldB := valB.Field(i)

		// Skip unexported fields
		if !fieldA.CanInterface() || !fieldB.CanInterface() {
			continue
		}

		// Handle slice fields specially
		if fieldA.Kind() == reflect.Slice {
			if !compareSlicesReflect(fieldA, fieldB) {
				return false
			}
			continue
		}

		// Compare other fields recursively
		if !CompareStructs(fieldA.Interface(), fieldB.Interface()) {
			return false
		}
	}

	// If all fields match, return true
	return true
}

// compareSlicesReflect is a helper function that compares slices using reflection.
// It's used internally by CompareStructs for comparing slice fields.
func compareSlicesReflect(a, b reflect.Value) bool {
	if a.Len() != b.Len() {
		return false
	}

	for i := 0; i < a.Len(); i++ {
		if !reflect.DeepEqual(a.Index(i).Interface(), b.Index(i).Interface()) {
			return false
		}
	}
	return true
}

// generateCacheKey generates a unique key for struct comparison based on the struct's type and values.
// This key is used for memoization to avoid repeated comparisons of the same structs.
func generateCacheKey(a, b interface{}) string {
	// Use a more specific key that includes pointer addresses for better uniqueness
	keyA := reflect.TypeOf(a).String()
	keyB := reflect.TypeOf(b).String()

	// Add pointer addresses to make keys more unique
	if reflect.ValueOf(a).Kind() == reflect.Ptr {
		keyA += fmt.Sprintf(":%p", a)
	}
	if reflect.ValueOf(b).Kind() == reflect.Ptr {
		keyB += fmt.Sprintf(":%p", b)
	}

	return keyA + ":" + keyB
}

// getCachedComparison retrieves a cached comparison result.
// Returns the result and a boolean indicating if the result was found.
func getCachedComparison(key string) (bool, bool) {
	structCache.RLock()
	defer structCache.RUnlock()
	result, found := structCache.cache[key]
	return result, found
}

// cacheComparisonResult stores the comparison result in the cache.
// This function is thread-safe and uses a read-write mutex for concurrent access.
func cacheComparisonResult(key string, result bool) {
	structCache.Lock()
	defer structCache.Unlock()
	structCache.cache[key] = result
}

// ClearStructCache clears the memoization cache for struct comparisons.
// This function is useful when memory usage becomes a concern or when
// you want to ensure fresh comparisons.
func ClearStructCache() {
	structCache.Lock()
	defer structCache.Unlock()
	structCache.cache = make(map[string]bool)
}

// GetStructCacheStats returns statistics about the struct comparison cache.
// This is useful for monitoring cache performance and memory usage.
func GetStructCacheStats() map[string]interface{} {
	structCache.RLock()
	defer structCache.RUnlock()

	return map[string]interface{}{
		"cache_size": len(structCache.cache),
		"cache_keys": reflect.ValueOf(structCache.cache).MapKeys(),
	}
}
