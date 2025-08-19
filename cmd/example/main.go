package main

import (
	"fmt"
	"time"

	"github.com/devrob-go/sliceutil/pkg/sliceutil"
)

func main() {
	fmt.Println("=== SliceUtil Examples ===")

	// Example 1: Basic slice comparison
	demonstrateSliceComparison()

	// Example 2: Finding differences
	demonstrateFindDifferences()

	// Example 3: Merge operations
	demonstrateMergeOperations()

	// Example 4: Statistical functions
	demonstrateStatisticalFunctions()

	// Example 5: Search and manipulation
	demonstrateSearchAndManipulation()

	// Example 6: Struct comparison
	demonstrateStructComparison()

	// Example 7: Error handling
	demonstrateErrorHandling()

	// Example 8: Performance features
	demonstratePerformanceFeatures()

	fmt.Println("\n=== All Examples Completed ===")
}

func demonstrateSliceComparison() {
	fmt.Println("1. Slice Comparison Examples:")
	fmt.Println("   ----------------------------")

	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	c := []int{5, 4, 3, 2, 1}

	// Basic comparison
	if sliceutil.CompareSlices(a, b) {
		fmt.Printf("   ✓ Slices a and b are equal: %v == %v\n", a, b)
	} else {
		fmt.Printf("   ✗ Slices a and b are different: %v != %v\n", a, b)
	}

	if sliceutil.CompareSlices(a, c) {
		fmt.Printf("   ✓ Slices a and c are equal: %v == %v\n", a, c)
	} else {
		fmt.Printf("   ✗ Slices a and c are different: %v != %v\n", a, c)
	}

	// Detailed comparison
	result := sliceutil.CompareSlicesWithResult(a, c)
	fmt.Printf("   Detailed comparison result: %s\n", result.Message)
	if !result.Equal {
		fmt.Printf("   Difference count: %d\n", result.Details["difference_count"])
	}

	fmt.Println()
}

func demonstrateFindDifferences() {
	fmt.Println("2. Finding Differences Examples:")
	fmt.Println("   ------------------------------")

	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}

	differences := sliceutil.FindDifferences(a, b)
	fmt.Printf("   Differences between %v and %v: %v\n", a, b, differences)

	// With counts
	differencesWithCount := sliceutil.FindDifferencesWithCount(a, b)
	fmt.Printf("   Differences with counts: %v\n", differencesWithCount)

	// String slices
	strA := []string{"apple", "banana", "cherry", "date"}
	strB := []string{"banana", "cherry", "elderberry", "fig"}
	strDifferences := sliceutil.FindDifferences(strA, strB)
	fmt.Printf("   String differences: %v\n", strDifferences)

	fmt.Println()
}

func demonstrateMergeOperations() {
	fmt.Println("3. Merge Operations Examples:")
	fmt.Println("   ---------------------------")

	a := []int{5, 1, 3}
	b := []int{4, 2, 6}

	// Merge with ascending order
	mergedAsc := sliceutil.MergeSlicesInt(a, b, sliceutil.OrderAsc)
	fmt.Printf("   Merged ascending: %v + %v = %v\n", a, b, mergedAsc)

	// Merge with descending order
	mergedDesc := sliceutil.MergeSlicesInt(a, b, sliceutil.OrderDesc)
	fmt.Printf("   Merged descending: %v + %v = %v\n", a, b, mergedDesc)

	// String merge
	strA := []string{"banana", "apple"}
	strB := []string{"cherry", "date"}
	mergedStr := sliceutil.MergeSlicesString(strA, strB, sliceutil.OrderAsc)
	fmt.Printf("   String merge: %v + %v = %v\n", strA, strB, mergedStr)

	// Generic merge with custom comparison
	type Person struct {
		Name string
		Age  int
	}
	peopleA := []Person{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}}
	peopleB := []Person{{Name: "Charlie", Age: 35}}
	less := func(a, b Person) bool { return a.Age < b.Age }
	mergedPeople := sliceutil.MergeSlicesGeneric(peopleA, peopleB, sliceutil.OrderAsc, less)
	fmt.Printf("   People merge (by age): %+v\n", mergedPeople)

	fmt.Println()
}

func demonstrateStatisticalFunctions() {
	fmt.Println("4. Statistical Functions Examples:")
	fmt.Println("   ---------------------------------")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Basic statistics
	max, err := sliceutil.MaxInt(slice)
	if err == nil {
		fmt.Printf("   Max value: %d\n", max)
	}

	min, err := sliceutil.MinInt(slice)
	if err == nil {
		fmt.Printf("   Min value: %d\n", min)
	}

	sum, err := sliceutil.SumInt(slice)
	if err == nil {
		fmt.Printf("   Sum: %d\n", sum)
	}

	avg, err := sliceutil.AverageInt(slice)
	if err == nil {
		fmt.Printf("   Average: %.2f\n", avg)
	}

	// Comprehensive statistics
	stats, err := sliceutil.GetSliceStats(slice)
	if err == nil {
		fmt.Printf("   Comprehensive stats:\n")
		fmt.Printf("     Length: %d\n", stats.Length)
		fmt.Printf("     Min: %v\n", stats.Min)
		fmt.Printf("     Max: %v\n", stats.Max)
		fmt.Printf("     Sum: %v\n", stats.Sum)
		fmt.Printf("     Average: %v\n", stats.Average)
		fmt.Printf("     Has Duplicates: %t\n", stats.HasDuplicates)
	}

	fmt.Println()
}

func demonstrateSearchAndManipulation() {
	fmt.Println("5. Search and Manipulation Examples:")
	fmt.Println("   ----------------------------------")

	slice := []int{1, 2, 3, 4, 5, 2, 3, 6}

	// Search operations
	contains := sliceutil.Contains(slice, 3)
	fmt.Printf("   Contains 3: %t\n", contains)

	index := sliceutil.IndexOf(slice, 3)
	fmt.Printf("   Index of 3: %d\n", index)

	count := sliceutil.CountOccurrences(slice, 2)
	fmt.Printf("   Count of 2: %d\n", count)

	// Manipulation operations
	reversed := sliceutil.ReverseCopy(slice)
	fmt.Printf("   Original: %v\n", slice)
	fmt.Printf("   Reversed copy: %v\n", reversed)

	unique := sliceutil.RemoveDuplicates(slice)
	fmt.Printf("   With duplicates: %v\n", slice)
	fmt.Printf("   Without duplicates: %v\n", unique)

	// Sorting check
	isSorted := sliceutil.IsSortedInt(slice)
	fmt.Printf("   Is sorted: %t\n", isSorted)

	fmt.Println()
}

func demonstrateStructComparison() {
	fmt.Println("6. Struct Comparison Examples:")
	fmt.Println("   ----------------------------")

	type Address struct {
		City   string
		Street string
		Zip    string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address
		Hobbies []string
	}

	person1 := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:   "New York",
			Street: "5th Ave",
			Zip:    "10001",
		},
		Hobbies: []string{"reading", "swimming"},
	}

	person2 := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:   "New York",
			Street: "5th Ave",
			Zip:    "10001",
		},
		Hobbies: []string{"reading", "swimming"},
	}

	person3 := Person{
		Name: "Bob",
		Age:  25,
		Address: Address{
			City:   "Los Angeles",
			Street: "Sunset Blvd",
			Zip:    "90210",
		},
		Hobbies: []string{"gaming", "cooking"},
	}

	// Compare identical structs
	if sliceutil.CompareStructs(person1, person2) {
		fmt.Printf("   ✓ Person1 and Person2 are identical\n")
	} else {
		fmt.Printf("   ✗ Person1 and Person2 are different\n")
	}

	// Compare different structs
	if sliceutil.CompareStructs(person1, person3) {
		fmt.Printf("   ✓ Person1 and Person3 are identical\n")
	} else {
		fmt.Printf("   ✗ Person1 and Person3 are different\n")
	}

	// Cache statistics
	stats := sliceutil.GetStructCacheStats()
	fmt.Printf("   Struct cache size: %d\n", stats["cache_size"])

	fmt.Println()
}

func demonstrateErrorHandling() {
	fmt.Println("7. Error Handling Examples:")
	fmt.Println("   -------------------------")

	// Test with nil slice
	_, err := sliceutil.MaxInt(nil)
	if err != nil {
		fmt.Printf("   ✓ Properly handled nil slice: %v\n", err)
	}

	// Test with empty slice
	_, err = sliceutil.MaxInt([]int{})
	if err != nil {
		fmt.Printf("   ✓ Properly handled empty slice: %v\n", err)
	}

	// Test with valid slice
	max, err := sliceutil.MaxInt([]int{1, 2, 3})
	if err == nil {
		fmt.Printf("   ✓ Successfully found max: %d\n", max)
	}

	// Test type mismatch in merge
	_, err = sliceutil.MergeSlices([]int{1, 2, 3}, []string{"a", "b"}, sliceutil.OrderAsc)
	if err != nil {
		fmt.Printf("   ✓ Properly handled type mismatch: %v\n", err)
	}

	fmt.Println()
}

func demonstratePerformanceFeatures() {
	fmt.Println("8. Performance Features Examples:")
	fmt.Println("   ------------------------------")

	// Large slice operations
	size := 10000
	largeSlice := make([]int, size)
	for i := 0; i < size; i++ {
		largeSlice[i] = i
	}

	// Performance test for comparison
	fmt.Printf("   Testing performance with slice of size %d...\n", size)

	// This should be fast due to early length check
	equal := sliceutil.CompareSlices(largeSlice, largeSlice)
	fmt.Printf("   Large slice comparison result: %t\n", equal)

	// Test struct cache performance
	fmt.Printf("   Testing struct comparison cache...\n")

	type TestStruct struct {
		ID   int
		Name string
		Data []int
	}

	// Create test structs
	testStruct := TestStruct{
		ID:   1,
		Name: "Test",
		Data: make([]int, 100),
	}

	// First comparison (populates cache)
	start := time.Now()
	sliceutil.CompareStructs(testStruct, testStruct)
	firstTime := time.Since(start)

	// Second comparison (uses cache)
	start = time.Now()
	sliceutil.CompareStructs(testStruct, testStruct)
	secondTime := time.Since(start)

	fmt.Printf("   First comparison: %v\n", firstTime)
	fmt.Printf("   Second comparison (cached): %v\n", secondTime)
	fmt.Printf("   Cache speedup: %.2fx\n", float64(firstTime)/float64(secondTime))

	// Clear cache
	sliceutil.ClearStructCache()
	stats := sliceutil.GetStructCacheStats()
	fmt.Printf("   Cache cleared, size: %d\n", stats["cache_size"])

	fmt.Println()
}
