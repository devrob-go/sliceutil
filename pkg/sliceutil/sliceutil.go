// Package sliceutil provides comprehensive utilities for working with Go slices.
// It includes functions for comparison, merging, finding differences, and other common slice operations.
//
// The package is designed with the following principles:
// - Type safety using Go generics where applicable
// - Proper error handling instead of panics
// - Comprehensive test coverage
// - Performance optimization for common operations
// - Clear and consistent API design
package sliceutil

import (
	"errors"
	"sync"
)

// Common errors that can be returned by sliceutil functions
var (
	ErrEmptySlice      = errors.New("slice cannot be empty")
	ErrNilSlice        = errors.New("slice cannot be nil")
	ErrTypeMismatch    = errors.New("slice types do not match")
	ErrUnsupportedType = errors.New("unsupported slice type")
)

// OrderType represents the sorting order for merge operations
type OrderType string

const (
	// OrderAsc sorts elements in ascending order
	OrderAsc OrderType = "ASC"
	// OrderDesc sorts elements in descending order
	OrderDesc OrderType = "DESC"
)

// Result represents the result of comparing two slices
type Result string

const (
	// ResultAGreater indicates slice A has a greater sum
	ResultAGreater Result = "a is greater"
	// ResultBGreater indicates slice B has a greater sum
	ResultBGreater Result = "b is greater"
	// ResultEqual indicates both slices have equal sums
	ResultEqual Result = "both are equal"
)

// CompareResult holds the result of a slice comparison operation
type CompareResult struct {
	Equal   bool
	Message string
	Details map[string]interface{}
}

// SliceStats provides statistical information about a slice
type SliceStats struct {
	Length        int
	Min           interface{}
	Max           interface{}
	Sum           interface{}
	Average       interface{}
	HasDuplicates bool
}

// Memoization cache for struct comparisons to improve performance
var structCache = struct {
	sync.RWMutex
	cache map[string]bool
}{cache: make(map[string]bool)}
