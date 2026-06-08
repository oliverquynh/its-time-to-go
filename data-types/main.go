package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Go has 4 signed integer types: int8, int16, int32 and int64.\n\n")

	fmt.Println("int8")
	fmt.Println("  - It's 1 byte (8 bits).")
	fmt.Printf("  - Range: min = %d, max = %d\n\n", math.MinInt8, math.MaxInt8)

	fmt.Println("int16")
	fmt.Println("  - It's 2 bytes (16 bits).")
	fmt.Printf("  - Range: min = %d, max = %d\n\n", math.MinInt16, math.MaxInt16)

	fmt.Println("int32")
	fmt.Println("  - It's 3 bytes (32 bits).")
	fmt.Printf("  - Range: min = %d, max = %d\n\n", math.MinInt32, math.MaxInt32)

	fmt.Println("int64")
	fmt.Println("  - It's 4 bytes (64 bits).")
	fmt.Printf("  - Range: min = %d, max = %d\n\n", math.MinInt64, math.MaxInt64)

	fmt.Printf("Go has 4 unsigned integer types: uint8, uint16, uint32 and uint64.\n\n")

	fmt.Println("uint8")
	fmt.Println("  - It's 1 byte (8 bits).")
	fmt.Printf("  - Range: min = 0, max = %d\n\n",  math.MaxUint8)

	fmt.Println("uint16")
	fmt.Println("  - It's 2 bytes (16 bits).")
	fmt.Printf("  - Range: min = 0, max = %d\n\n",  math.MaxUint16)

	fmt.Println("uint32")
	fmt.Println("  - It's 3 bytes (32 bits).")
	fmt.Printf("  - Range: min = 0, max = %d\n\n",  math.MaxUint32)

	// TODO: Why cannot %d represent math.MaxUint64
	// fmt.Println("int64")
	// fmt.Println("  - It's 4 bytes (64 bits).")
	// fmt.Printf("  - Range: min = 0, max = %d\n\n",  math.MaxUint64)

	fmt.Printf("Go has 2 floating point types: float32 and float64.\n\n")

	fmt.Println("float32")
	fmt.Println("  - It's 3 bytes (32 bits).")
	fmt.Printf("  - Range: min = %f, max = %f\n\n", math.SmallestNonzeroFloat32, math.MaxFloat32)

	fmt.Println("float64")
	fmt.Println("  - It's 8 bytes (64 bits).")
	fmt.Printf("  - Range: min = %f, max = %f\n\n", math.SmallestNonzeroFloat64, math.MaxFloat64)
}
