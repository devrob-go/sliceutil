package sliceutil

import (
	"reflect"
	"sync"
)

// Memoization cache for CompareStructs
var structCache = struct {
	sync.RWMutex
	cache map[string]bool
}{cache: make(map[string]bool)}

// CompareStructs compares two structs deeply, with memoization to avoid redundant comparisons
func CompareStructs(a, b interface{}) bool {
	// Check if both are nil or have different types early on.
	if a == nil || b == nil {
		return a == b
	}

	// If the types are different, they cannot be equal.
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	// Generate a key for caching based on the struct's type and data
	cacheKey := generateCacheKey(a, b)
	structCache.RLock()
	if cachedResult, found := structCache.cache[cacheKey]; found {
		structCache.RUnlock()
		return cachedResult
	}
	structCache.RUnlock()

	// Use reflection to inspect the struct fields
	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)

	// Ensure that we can compare field values
	if valA.Kind() == reflect.Ptr {
		valA = valA.Elem()
	}
	if valB.Kind() == reflect.Ptr {
		valB = valB.Elem()
	}

	// Iterate over each field of the struct and compare the values
	for i := 0; i < valA.NumField(); i++ {
		fieldA := valA.Field(i)
		fieldB := valB.Field(i)

		// Skip unexported fields (i.e., those with lowercase starting letters)
		if fieldA.CanInterface() && fieldB.CanInterface() {
			if fieldA.Kind() == reflect.Struct {
				// Recursively compare struct fields
				if !CompareStructs(fieldA.Interface(), fieldB.Interface()) {
					cacheComparisonResult(cacheKey, false)
					return false
				}
			} else if fieldA.Kind() == reflect.Slice && fieldB.Kind() == reflect.Slice {
				// If both fields are slices, compare them using CompareSlices
				// Assert the types of the slices before passing them to CompareSlices
				switch fieldA.Type().Elem().Kind() {
				case reflect.Int:
					// Assert slices as []int
					aSlice := fieldA.Interface().([]int)
					bSlice := fieldB.Interface().([]int)
					if !CompareSlices(aSlice, bSlice) {
						cacheComparisonResult(cacheKey, false)
						return false
					}
				case reflect.String:
					// Assert slices as []string
					aSlice := fieldA.Interface().([]string)
					bSlice := fieldB.Interface().([]string)
					if !CompareSlices(aSlice, bSlice) {
						cacheComparisonResult(cacheKey, false)
						return false
					}
				default:
					// Return false if the slice types are not supported
					cacheComparisonResult(cacheKey, false)
					return false
				}
			} else if fieldA.Interface() != fieldB.Interface() {
				// If the field values differ, return false
				cacheComparisonResult(cacheKey, false)
				return false
			}
		}
	}

	// Cache and return the result if structs match
	cacheComparisonResult(cacheKey, true)
	return true
}

// generateCacheKey generates a unique key for struct comparison based on the struct's type and values.
func generateCacheKey(a, b interface{}) string {
	return reflect.TypeOf(a).String() + reflect.TypeOf(b).String()
}

// cacheComparisonResult stores the comparison result in the cache.
func cacheComparisonResult(key string, result bool) {
	structCache.Lock()
	defer structCache.Unlock()
	structCache.cache[key] = result
}
