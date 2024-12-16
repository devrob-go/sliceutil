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

// CompareStructs compares two structs deeply, supporting nested structs and pointers.
func CompareStructs(a, b interface{}) bool {
	// If both are nil, they are equal.
	if a == nil && b == nil {
		return true
	}

	// If one is nil, they are not equal.
	if a == nil || b == nil {
		return false
	}

	// Check if the types are the same.
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	// Use reflection to inspect the struct fields.
	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)

	// Dereference pointers, if applicable.
	if valA.Kind() == reflect.Ptr {
		if valA.IsNil() || valB.IsNil() {
			return valA.IsNil() && valB.IsNil()
		}
		valA = valA.Elem()
		valB = valB.Elem()
	}

	// Check if both are structs.
	if valA.Kind() != reflect.Struct || valB.Kind() != reflect.Struct {
		return reflect.DeepEqual(a, b)
	}

	// Compare each field of the structs.
	for i := 0; i < valA.NumField(); i++ {
		fieldA := valA.Field(i)
		fieldB := valB.Field(i)

		// Skip unexported fields.
		if !fieldA.CanInterface() || !fieldB.CanInterface() {
			continue
		}

		// Compare fields recursively.
		if !CompareStructs(fieldA.Interface(), fieldB.Interface()) {
			return false
		}
	}

	// If all fields match, return true.
	return true
}

// Helper: Compare slices using reflection.
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
func generateCacheKey(a, b interface{}) string {
	return reflect.TypeOf(a).String() + reflect.TypeOf(b).String()
}

// cacheComparisonResult stores the comparison result in the cache.
func cacheComparisonResult(key string, result bool) {
	structCache.Lock()
	defer structCache.Unlock()
	structCache.cache[key] = result
}
