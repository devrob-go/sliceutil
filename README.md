# SliceUtil - Go Package for Slice and Struct Operations

## Overview

`SliceUtil` is a Go package designed to provide various utilities and helper functions for working with slices and structs. It includes functions to compare slices, find differences between slices, merge slices with sorting options, and perform deep comparisons on structs, including nested structs and slices within structs.

The package provides the following functionalities:

- **CompareSlices**: Compares two slices of any comparable type (e.g., `int`, `string`) for equality.
- **FindDifferences**: Finds and returns the unique elements between two slices.
- **MergeSlices**: Merges two slices and sorts the resulting slice in ascending or descending order.
- **MaxIntSlice**: Returns the maximum value from a slice of integers.
- **MinIntSlice**: Returns the minimum value from a slice of integers.
- **SumSlice**: Sums up the elements of a slice of integers or floats.
- **CompareStructs**: Compares two structs deeply using reflection.
- **CompareMultiLayerStructs**: Compares deeply nested structs, including slices within structs.
- **Memoization**: Implements caching to optimize redundant struct comparisons.

## Features

- **Generics Support**: Uses Go's generics (type parameters) to support a wide variety of types for slices (e.g., `int`, `string`).
- **Memoization**: Struct comparisons are cached for efficiency, especially when comparing large or repetitive structs.
- **Deep Comparison**: Structs, slices, and nested fields are compared recursively.
- **Edge Cases**: Handles `nil` values, slice and struct type mismatches, and deeply nested structures.

## Installation

To use the package in your project, install it via `go get`:

```bash
go get github.com/yourusername/sliceutil
```

## Functions
`CompareSlices[T comparable](a []T, b []T) bool`
Compares two slices for equality (both in length and values). Returns true if the slices are identical, otherwise false.

`FindDifferences[T comparable](a []T, b []T) []T`
Finds and returns a slice of elements that are unique to each of the two slices. The output is a new slice containing the unique elements from both slices.

`MergeSlices[T comparable](a []T, b []T, asc bool) []T`
Merges two slices into one and sorts the result. The asc flag determines the sort order: true for ascending (A-Z, smaller to bigger), false for descending (Z-A, bigger to smaller).

`MaxIntSlice(a []int) int`
Returns the maximum value in an integer slice. If the slice is empty, it returns math.MinInt.

`MinIntSlice(a []int) int`
Returns the minimum value in an integer slice. If the slice is empty, it returns math.MaxInt.

`SumSlice[T int | float64](a []T) T`
Calculates and returns the sum of elements in a slice. Supports both int and float64 types.

`CompareStructs(a, b interface{}) bool`
Compares two structs deeply using reflection. Returns true if the structs are identical, otherwise false. Supports memoization to avoid redundant comparisons.

`CompareMultiLayerStructs(a, b interface{}) bool`
Compares structs with multiple layers of nesting, including slices within structs. Utilizes recursion to compare deeply nested structs and slices.

# Example Usage
## Compare Slices

```go
package main

import (
	"fmt"
	"github.com/yourusername/sliceutil"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	if sliceutil.CompareSlices(a, b) {
		fmt.Println("Slices are equal")
	} else {
		fmt.Println("Slices are different")
	}
}

```

## Find Differences Between Slices

```go
package main

import (
	"fmt"
	"github.com/yourusername/sliceutil"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	diff := sliceutil.FindDifferences(a, b)
	fmt.Println("Differences:", diff)
}

```

## Merge and Sort Slices

```go
package main

import (
	"fmt"
	"github.com/yourusername/sliceutil"
)

func main() {
	a := []int{3, 1}
	b := []int{4, 2}

	merged := sliceutil.MergeSlices(a, b, true) // true for ascending order
	fmt.Println("Merged and Sorted:", merged)
}

```

## Compare Structs

```go
package main

import (
	"fmt"
	"github.com/yourusername/sliceutil"
)

type Person struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	a := Person{"Alice", 30, true}
	b := Person{"Alice", 30, true}

	if sliceutil.CompareStructs(a, b) {
		fmt.Println("Structs are equal")
	} else {
		fmt.Println("Structs are different")
	}
}

```

## Compare Multi-Layer Structs

```go
package main

import (
	"fmt"
	"github.com/yourusername/sliceutil"
)

type Address struct {
	City   string
	Street string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	a := Person{"Alice", 30, Address{"New York", "5th Ave"}}
	b := Person{"Alice", 30, Address{"New York", "5th Ave"}}

	if sliceutil.CompareMultiLayerStructs(a, b) {
		fmt.Println("Multi-layer structs are equal")
	} else {
		fmt.Println("Multi-layer structs are different")
	}
}

```
# Tests

The package includes unit tests for all functions. To run the tests:
```bash
go test -v ./...
```

# License

MIT License. See LICENSE for more details.
```bash
You can copy and paste the entire content of this `README.md` file into your project.
```