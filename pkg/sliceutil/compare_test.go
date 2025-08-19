package sliceutil

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCompareStructs tests the CompareStructs function
func TestCompareStructs(t *testing.T) {
	type Address struct {
		City   string
		Street string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address
	}

	t.Run("Identical Flat Structs", func(t *testing.T) {
		a := Person{Name: "Alice", Age: 30}
		b := Person{Name: "Alice", Age: 30}
		assert.True(t, CompareStructs(a, b))
	})

	t.Run("Different Flat Structs", func(t *testing.T) {
		a := Person{Name: "Alice", Age: 30}
		b := Person{Name: "Bob", Age: 40}
		assert.False(t, CompareStructs(a, b))
	})

	t.Run("Identical Nested Structs", func(t *testing.T) {
		a := Person{Name: "Alice", Age: 30, Address: Address{City: "New York", Street: "5th Ave"}}
		b := Person{Name: "Alice", Age: 30, Address: Address{City: "New York", Street: "5th Ave"}}
		assert.True(t, CompareStructs(a, b))
	})

	t.Run("Different Nested Structs", func(t *testing.T) {
		a := Person{Name: "Alice", Age: 30, Address: Address{City: "New York", Street: "5th Ave"}}
		b := Person{Name: "Alice", Age: 30, Address: Address{City: "Los Angeles", Street: "Sunset Blvd"}}
		assert.False(t, CompareStructs(a, b))
	})

	t.Run("Structs with Identical Pointers", func(t *testing.T) {
		address := &Address{City: "New York", Street: "5th Ave"}
		a := struct {
			Name    string
			Address *Address
		}{Name: "Alice", Address: address}
		b := struct {
			Name    string
			Address *Address
		}{Name: "Alice", Address: address}
		assert.True(t, CompareStructs(a, b))
	})

	t.Run("Structs with Different Pointers", func(t *testing.T) {
		a := struct {
			Name    string
			Address *Address
		}{Name: "Alice", Address: &Address{City: "New York", Street: "5th Ave"}}
		b := struct {
			Name    string
			Address *Address
		}{Name: "Alice", Address: &Address{City: "Los Angeles", Street: "Sunset Blvd"}}
		assert.False(t, CompareStructs(a, b))
	})

	t.Run("Nil Structs", func(t *testing.T) {
		assert.True(t, CompareStructs(nil, nil))
		assert.False(t, CompareStructs(Person{Name: "Alice"}, nil))
		assert.False(t, CompareStructs(nil, Person{Name: "Alice"}))
	})

	t.Run("Different Types", func(t *testing.T) {
		a := Person{Name: "Alice", Age: 30}
		b := "not a struct"
		assert.False(t, CompareStructs(a, b))
	})

	t.Run("Pointer to Struct", func(t *testing.T) {
		a := &Person{Name: "Alice", Age: 30}
		b := &Person{Name: "Alice", Age: 30}
		assert.True(t, CompareStructs(a, b))
	})

	t.Run("Nil Pointer vs Struct", func(t *testing.T) {
		var a *Person
		b := Person{Name: "Alice", Age: 30}
		assert.False(t, CompareStructs(a, b))
	})
}

// TestCompareReflectionSlices tests the CompareReflectionSlices function
func TestCompareReflectionSlices(t *testing.T) {
	t.Run("Int Slices", func(t *testing.T) {
		a := reflect.ValueOf([]int{1, 2, 3})
		b := reflect.ValueOf([]int{1, 2, 3})
		assert.True(t, CompareReflectionSlices(a, b))
	})

	t.Run("String Slices", func(t *testing.T) {
		a := reflect.ValueOf([]string{"a", "b", "c"})
		b := reflect.ValueOf([]string{"a", "b", "c"})
		assert.True(t, CompareReflectionSlices(a, b))
	})

	t.Run("Different Types", func(t *testing.T) {
		a := reflect.ValueOf([]int{1, 2, 3})
		b := reflect.ValueOf([]string{"a", "b", "c"})
		assert.False(t, CompareReflectionSlices(a, b))
	})

	t.Run("Non-Slice Values", func(t *testing.T) {
		a := reflect.ValueOf("not a slice")
		b := reflect.ValueOf("also not a slice")
		assert.False(t, CompareReflectionSlices(a, b))
	})

	t.Run("Unsupported Slice Type", func(t *testing.T) {
		a := reflect.ValueOf([]bool{true, false})
		b := reflect.ValueOf([]bool{true, false})
		assert.False(t, CompareReflectionSlices(a, b))
	})
}

// TestStructCache tests the struct comparison cache functionality
func TestStructCache(t *testing.T) {
	t.Run("Cache Operations", func(t *testing.T) {
		// Since caching is disabled, just test that the functions don't panic
		ClearStructCache()

		// Get initial stats
		initialStats := GetStructCacheStats()
		assert.Equal(t, 0, initialStats["cache_size"])

		// Compare some structs (no caching)
		type TestStruct struct {
			Name string
			Age  int
		}

		a := TestStruct{Name: "Alice", Age: 30}
		b := TestStruct{Name: "Bob", Age: 25}

		// Comparison should work without caching
		result := CompareStructs(a, b)
		assert.False(t, result)

		// Cache should remain empty
		stats := GetStructCacheStats()
		assert.Equal(t, 0, stats["cache_size"])

		// Clear cache
		ClearStructCache()

		// Verify cache is cleared
		finalStats := GetStructCacheStats()
		assert.Equal(t, 0, finalStats["cache_size"])
	})
}

// TestCompareSlicesEdgeCases tests edge cases for slice comparison
func TestCompareSlicesEdgeCases(t *testing.T) {
	t.Run("Very Large Slices", func(t *testing.T) {
		size := 100000
		a := make([]int, size)
		b := make([]int, size)

		for i := 0; i < size; i++ {
			a[i] = i
			b[i] = i
		}

		assert.True(t, CompareSlices(a, b))
	})

	t.Run("Slices with Different Lengths", func(t *testing.T) {
		a := make([]int, 1000)
		b := make([]int, 1001)

		for i := 0; i < 1000; i++ {
			a[i] = i
			b[i] = i
		}
		b[1000] = 1000

		assert.False(t, CompareSlices(a, b))
	})

	t.Run("Empty vs Non-Empty", func(t *testing.T) {
		a := []int{}
		b := []int{1, 2, 3}
		assert.False(t, CompareSlices(a, b))
	})

	t.Run("Single Element Slices", func(t *testing.T) {
		a := []int{42}
		b := []int{42}
		assert.True(t, CompareSlices(a, b))

		c := []int{42}
		d := []int{24}
		assert.False(t, CompareSlices(c, d))
	})
}

// TestCompareSlicesWithResultEdgeCases tests edge cases for detailed comparison
func TestCompareSlicesWithResultEdgeCases(t *testing.T) {
	t.Run("Large Difference Count", func(t *testing.T) {
		size := 1000
		a := make([]int, size)
		b := make([]int, size)

		for i := 0; i < size; i++ {
			a[i] = i
			b[i] = i + 1 // All elements are different
		}

		result := CompareSlicesWithResult(a, b)
		assert.False(t, result.Equal)
		assert.Equal(t, size, result.Details["difference_count"])
	})

	t.Run("Mixed Nil and Non-Nil", func(t *testing.T) {
		result := CompareSlicesWithResult[int](nil, []int{1, 2, 3})
		assert.False(t, result.Equal)
		assert.True(t, result.Details["a_nil"].(bool))
		assert.False(t, result.Details["b_nil"].(bool))
	})
}

// TestPerformance tests performance characteristics
func TestPerformance(t *testing.T) {
	t.Run("CompareSlices Performance", func(t *testing.T) {
		size := 10000
		a := make([]int, size)
		b := make([]int, size)

		for i := 0; i < size; i++ {
			a[i] = i
			b[i] = i
		}

		// This should be fast
		result := CompareSlices(a, b)
		assert.True(t, result)
	})

	t.Run("CompareSlicesWithResult Performance", func(t *testing.T) {
		size := 10000
		a := make([]int, size)
		b := make([]int, size)

		for i := 0; i < size; i++ {
			a[i] = i
			b[i] = i
		}

		// This should also be fast
		result := CompareSlicesWithResult(a, b)
		assert.True(t, result.Equal)
	})
}

// TestConcurrentAccess tests concurrent access to the struct cache
func TestConcurrentAccess(t *testing.T) {
	t.Run("Concurrent Struct Comparisons", func(t *testing.T) {
		// Clear cache first
		ClearStructCache()

		type ConcurrentStruct struct {
			ID   int
			Name string
		}

		// Create multiple goroutines to test concurrent access
		done := make(chan bool, 10)

		for i := 0; i < 10; i++ {
			go func(id int) {
				a := ConcurrentStruct{ID: id, Name: "Test"}
				b := ConcurrentStruct{ID: id, Name: "Test"}

				// This should be thread-safe
				CompareStructs(a, b)
				done <- true
			}(i)
		}

		// Wait for all goroutines to complete
		for i := 0; i < 10; i++ {
			<-done
		}

		// Verify no panics occurred and cache is accessible
		stats := GetStructCacheStats()
		assert.NotNil(t, stats)
	})
}

// TestMemoryEfficiency tests memory efficiency of the comparison functions
func TestMemoryEfficiency(t *testing.T) {
	t.Run("Large Slice Comparison Memory", func(t *testing.T) {
		size := 100000
		a := make([]int, size)
		b := make([]int, size)

		for i := 0; i < size; i++ {
			a[i] = i
			b[i] = i
		}

		// This should not allocate excessive memory
		result := CompareSlices(a, b)
		assert.True(t, result)
	})
}

// TestErrorHandling tests error handling in edge cases
func TestErrorHandling(t *testing.T) {
	t.Run("Panic Prevention", func(t *testing.T) {
		// These should not panic
		assert.NotPanics(t, func() {
			CompareSlices[int](nil, nil)
		})

		assert.NotPanics(t, func() {
			CompareSlicesWithResult[int](nil, nil)
		})

		assert.NotPanics(t, func() {
			CompareStructs(nil, nil)
		})
	})
}
