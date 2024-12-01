package main

// basic types: numbers, strings, booleans
// aggregate types: array, struct
// reference types: pointers, slices, maps, functions, channels
// interface type

import (
	"myapp/types"
)

func main() {
	// basic
	types.Numerics()
	types.Strings()
	types.Booleans()

	// aggregate
	types.Arrays() // not used as much, due to slices
	types.Structs()

	// reference
	types.Pointers()
	types.Slices()
	types.Maps()
	types.Functions()
	types.Channels()

	// interface
	types.Interfaces()
}
