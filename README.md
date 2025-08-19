# SliceUtil

A production-grade Go package providing comprehensive utilities for working with Go slices. Built with modern Go features including generics, proper error handling, and extensive test coverage.

## Features

- **Type Safety**: Uses Go generics for compile-time type safety
- **Comprehensive Operations**: Comparison, merging, finding differences, statistics, and more
- **Error Handling**: Proper error handling instead of panics
- **Performance Optimized**: Efficient algorithms with memoization for struct comparisons
- **Thread Safe**: Concurrent access support with proper synchronization
- **Well Tested**: Extensive test coverage with benchmarks
- **Production Ready**: Industry-standard code structure and documentation

## Installation

```bash
go get github.com/devrob-go/sliceutil
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/devrob-go/sliceutil"
)

func main() {
    // Compare slices
    a := []int{1, 2, 3, 4, 5}
    b := []int{1, 2, 3, 4, 5}
    
    if sliceutil.CompareSlices(a, b) {
        fmt.Println("Slices are equal!")
    }
    
    // Find differences
    c := []int{1, 2, 3, 6, 7}
    differences := sliceutil.FindDifferences(a, c)
    fmt.Printf("Differences: %v\n", differences) // [4, 5, 6, 7]
    
    // Merge and sort
    merged := sliceutil.MergeSlicesInt(a, c, sliceutil.OrderAsc)
    fmt.Printf("Merged: %v\n", merged) // [1, 1, 2, 2, 3, 3, 4, 5, 6, 7]
    
    // Get statistics
    stats, err := sliceutil.GetSliceStats(a)
    if err == nil {
        fmt.Printf("Length: %d, Min: %v, Max: %v, Sum: %v, Average: %.2f\n",
            stats.Length, stats.Min, stats.Max, stats.Sum, stats.Average)
    }
}
```

## API Reference

### Comparison Functions

#### `CompareSlices[T comparable](a, b []T) bool`
Compares two slices for equality in values and order.

```go
a := []int{1, 2, 3}
b := []int{1, 2, 3}
equal := sliceutil.CompareSlices(a, b) // true
```

#### `CompareSlicesWithResult[T comparable](a, b []T) CompareResult`
Provides detailed comparison results including difference locations.

```go
result := sliceutil.CompareSlicesWithResult(a, b)
if !result.Equal {
    fmt.Printf("Slices differ: %s\n", result.Message)
    fmt.Printf("Difference count: %d\n", result.Details["difference_count"])
}
```

#### `CompareStructs(a, b interface{}) bool`
Deep comparison of structs with memoization for performance.

```go
type Person struct {
    Name string
    Age  int
}

a := Person{Name: "Alice", Age: 30}
b := Person{Name: "Alice", Age: 30}
equal := sliceutil.CompareStructs(a, b) // true
```

### Utility Functions

#### `FindDifferences[T comparable](a, b []T) []T`
Returns unique values from both slices that are not in the other.

```go
a := []int{1, 2, 3, 4}
b := []int{3, 4, 5, 6}
differences := sliceutil.FindDifferences(a, b) // [1, 2, 5, 6]
```

#### `MaxInt(a []int) (int, error)`
Finds the maximum value in an int slice.

```go
max, err := sliceutil.MaxInt([]int{1, 5, 3, 9, 2})
if err == nil {
    fmt.Printf("Max: %d\n", max) // 9
}
```

#### `MinInt(a []int) (int, error)`
Finds the minimum value in an int slice.

```go
min, err := sliceutil.MinInt([]int{1, 5, 3, 9, 2})
if err == nil {
    fmt.Printf("Min: %d\n", min) // 1
}
```

#### `SumInt(a []int) (int, error)`
Calculates the sum of all integers in a slice.

```go
sum, err := sliceutil.SumInt([]int{1, 2, 3, 4, 5})
if err == nil {
    fmt.Printf("Sum: %d\n", sum) // 15
}
```

#### `AverageInt(a []int) (float64, error)`
Calculates the average of all integers in a slice.

```go
avg, err := sliceutil.AverageInt([]int{1, 2, 3, 4, 5})
if err == nil {
    fmt.Printf("Average: %.2f\n", avg) // 3.00
}
```

### Merge Functions

#### `MergeSlicesInt(a, b []int, order OrderType) []int`
Merges two int slices with specified sorting order.

```go
a := []int{5, 1, 3}
b := []int{4, 2, 6}
merged := sliceutil.MergeSlicesInt(a, b, sliceutil.OrderAsc)
// Result: [1, 2, 3, 4, 5, 6]
```

#### `MergeSlicesString(a, b []string, order OrderType) []string`
Merges two string slices with specified sorting order.

```go
a := []string{"banana", "apple"}
b := []string{"cherry", "date"}
merged := sliceutil.MergeSlicesString(a, b, sliceutil.OrderAsc)
// Result: ["apple", "banana", "cherry", "date"]
```

#### `MergeSlicesGeneric[T any](a, b []T, order OrderType, less func(T, T) bool) []T`
Generic merge function with custom comparison logic.

```go
type Person struct {
    Name string
    Age  int
}

a := []Person{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}}
b := []Person{{Name: "Charlie", Age: 35}}

less := func(a, b Person) bool { return a.Age < b.Age }
merged := sliceutil.MergeSlicesGeneric(a, b, sliceutil.OrderAsc, less)
```

### Search and Manipulation Functions

#### `Contains[T comparable](a []T, element T) bool`
Checks if a slice contains a specific element.

```go
slice := []int{1, 2, 3, 4, 5}
contains := sliceutil.Contains(slice, 3) // true
```

#### `IndexOf[T comparable](a []T, element T) int`
Returns the index of the first occurrence of an element.

```go
slice := []int{1, 2, 3, 4, 5}
index := sliceutil.IndexOf(slice, 3) // 2
```

#### `RemoveDuplicates[T comparable](a []T) []T`
Removes duplicate elements while preserving order.

```go
slice := []int{1, 2, 2, 3, 3, 4}
unique := sliceutil.RemoveDuplicates(slice) // [1, 2, 3, 4]
```

#### `Reverse[T any](a []T)`
Reverses the order of elements in a slice (modifies original).

```go
slice := []int{1, 2, 3, 4, 5}
sliceutil.Reverse(slice)
// slice is now [5, 4, 3, 2, 1]
```

#### `ReverseCopy[T any](a []T) []T`
Creates a reversed copy without modifying the original.

```go
original := []int{1, 2, 3, 4, 5}
reversed := sliceutil.ReverseCopy(original)
// original is unchanged, reversed is [5, 4, 3, 2, 1]
```

### Statistics Functions

#### `GetSliceStats(a []int) (SliceStats, error)`
Provides comprehensive statistical information about a slice.

```go
stats, err := sliceutil.GetSliceStats([]int{1, 2, 3, 4, 5})
if err == nil {
    fmt.Printf("Length: %d\n", stats.Length)
    fmt.Printf("Min: %v\n", stats.Min)
    fmt.Printf("Max: %v\n", stats.Max)
    fmt.Printf("Sum: %v\n", stats.Sum)
    fmt.Printf("Average: %v\n", stats.Average)
    fmt.Printf("Has Duplicates: %t\n", stats.HasDuplicates)
}
```

### Cache Management

#### `ClearStructCache()`
Clears the memoization cache for struct comparisons.

```go
sliceutil.ClearStructCache()
```

#### `GetStructCacheStats() map[string]interface{}`
Returns statistics about the struct comparison cache.

```go
stats := sliceutil.GetStructCacheStats()
fmt.Printf("Cache size: %d\n", stats["cache_size"])
```

## Error Handling

The package uses proper error handling instead of panics. Common errors include:

- `ErrEmptySlice`: Returned when a slice is empty but cannot be
- `ErrNilSlice`: Returned when a slice is nil but cannot be
- `ErrTypeMismatch`: Returned when slice types don't match
- `ErrUnsupportedType`: Returned when a type is not supported

```go
max, err := sliceutil.MaxInt([]int{})
if err != nil {
    switch {
    case errors.Is(err, sliceutil.ErrEmptySlice):
        fmt.Println("Cannot find max in empty slice")
    case errors.Is(err, sliceutil.ErrNilSlice):
        fmt.Println("Cannot find max in nil slice")
    }
}
```

## Performance Considerations

- **Slice Comparison**: O(n) time complexity for basic comparison
- **Struct Comparison**: Uses memoization to avoid repeated comparisons
- **Merge Operations**: O((n + m) * log(n + m)) time complexity due to sorting
- **Memory Usage**: Efficient memory usage with minimal allocations

## Thread Safety

- **Struct Comparison Cache**: Thread-safe with read-write mutex
- **Slice Operations**: All slice operations are thread-safe
- **Concurrent Access**: Safe for concurrent use in multiple goroutines

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Run benchmarks:

```bash
go test -bench=. ./...
```

## Examples

See the `examples/` directory for more detailed usage examples.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Version History

- **v1.0.0**: Initial release with comprehensive slice utilities
- Comprehensive comparison functions
- Merge and sort operations
- Statistical analysis functions
- Thread-safe struct comparison with memoization
- Extensive test coverage

## Support

For questions, issues, or contributions, please open an issue on GitHub or submit a pull request.