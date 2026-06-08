package main

import (
	"fmt"
	"time"
)

var aa int8
var bb int8 = 11
// cc := 22 // Shorthand is not allowed here.

// Single untyped constant
const Pi = 3.14159

// Multiple untyped constants
const (
	StatusOk = 200
	StatusBadRequest = 400
	StatusInternalServerError = 500
)

const (
    Timeout time.Duration = 30 * time.Second
    MaxSize int          = 1024
    Debug   bool         = false
)

// iota starts from 0, increased 1 by 1, reset per const block.
const (
	A = iota
	B // 1
	C // 2
)

const (
	AA = iota * iota // 0 * 0 = 0
	BB // 1 * 1 = 1
	CC // 2 * 2 = 4
)

func main() {
	fmt.Println("=== Variables ===")

	// Declare a variable without assigning a value.
	// Go will assign a default zero value based on the variable type.
	// This case, it'll be 0
	var a int8

	fmt.Printf("a is %d.\n", a)

	// Declare a variable with assigning a value.
	var b int8 = 1

	fmt.Printf("b is %d.\n", b)

	// The short hand to declare a variable and assign a value.
	// Go implicitly infers c is int.
	// only available inside functions
	c := 2

	fmt.Printf("c is %d.\n", c)

	fmt.Println("=== Constants ===")
	fmt.Printf("-> Pi = %f\n", Pi)
	fmt.Printf("-> StatusOk = %d\n", StatusOk)
	fmt.Printf("-> StatusBadRequest = %d\n", StatusBadRequest)
	fmt.Printf("-> StatusInternalServerError = %d\n", StatusInternalServerError)
	fmt.Printf("-> Timeout = %d\n", Timeout)
	fmt.Printf("-> MaxSide = %d\n", MaxSize)
	fmt.Printf("-> Debug = %t\n", Debug)

	fmt.Println("=== Iota ===")
	fmt.Printf("-> A is %d.\n", A)
	fmt.Printf("-> B is %d.\n", B)
	fmt.Printf("-> C is %d.\n", C)

	fmt.Printf("-> AA is %d.\n", AA)
	fmt.Printf("-> BB is %d.\n", BB)
	fmt.Printf("-> CC is %d.\n", CC)

}
