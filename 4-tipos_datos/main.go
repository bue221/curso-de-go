package main

import (
	"fmt"
	"math"
)

func main() {

	// Numericos
	// Int8, Int16, Int32, Int64, Int
	// Uint8, Uint16, Uint32, Uint64, Uint
	// Float32, Float64
	// Complex64, Complex128

	// Booleanos
	// Bool

	// Texto
	// String

	// Rune

	// Array
	// Slice
	// Map
	// Struct

	var name string = "Andres"
	var age int = 20
	var isStudent bool = true

	fmt.Println(name, age, isStudent)

	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxUint8)
	fmt.Println(math.MaxUint16)
	fmt.Println(math.MaxUint32)
}
