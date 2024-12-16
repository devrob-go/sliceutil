package sliceutil_test

import (
	"testing"

	"github.com/devrob-go/sliceutil"
	"github.com/stretchr/testify/assert"
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

	// Test case 1: Identical flat structs
	t.Run("Identical Flat Structs", func(t *testing.T) {
		a := Person{Name: "Alice", Age: 30}
		b := Person{Name: "Alice", Age: 30}
		assert.True(t, sliceutil.CompareStructs(a, b), "Structs should be identical")
	})

	// Test case 2: Different flat structs
	t.Run("Different Flat Structs", func(t *testing.T) {
		a := Person{Name: "Alice", Age: 30}
		b := Person{Name: "Bob", Age: 40}
		assert.False(t, sliceutil.CompareStructs(a, b), "Structs should be different")
	})

	// Test case 3: Identical nested structs
	t.Run("Identical Nested Structs", func(t *testing.T) {
		a := Person{Name: "Alice", Age: 30, Address: Address{City: "New York", Street: "5th Ave"}}
		b := Person{Name: "Alice", Age: 30, Address: Address{City: "New York", Street: "5th Ave"}}
		assert.True(t, sliceutil.CompareStructs(a, b), "Nested structs should be identical")
	})

	// Test case 4: Different nested structs
	t.Run("Different Nested Structs", func(t *testing.T) {
		a := Person{Name: "Alice", Age: 30, Address: Address{City: "New York", Street: "5th Ave"}}
		b := Person{Name: "Alice", Age: 30, Address: Address{City: "Los Angeles", Street: "Sunset Blvd"}}
		assert.False(t, sliceutil.CompareStructs(a, b), "Nested structs should be different")
	})

	// Test case 5: Structs with identical pointers
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
		assert.True(t, sliceutil.CompareStructs(a, b), "Structs with identical pointers should be equal")
	})

	// Test case 6: Structs with different pointers
	t.Run("Structs with Different Pointers", func(t *testing.T) {
		a := struct {
			Name    string
			Address *Address
		}{Name: "Alice", Address: &Address{City: "New York", Street: "5th Ave"}}
		b := struct {
			Name    string
			Address *Address
		}{Name: "Alice", Address: &Address{City: "Los Angeles", Street: "Sunset Blvd"}}
		assert.False(t, sliceutil.CompareStructs(a, b), "Structs with different pointers should not be equal")
	})
}
