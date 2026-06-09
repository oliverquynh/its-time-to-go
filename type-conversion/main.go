package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// (1) number -> number

	var a int32 = 123
	var b int64 = int64(a)
	var c float32 = float32(b)

	fmt.Printf("a = %d, b = %d, c = %f\n", a, b, c)

	// (2) string -> number

	var value = "123"

	intValue, err := strconv.Atoi(value)

	if err != nil {
		log.Fatalf("Conversion failed: %v", err)
		return
	}

	fmt.Printf("Value: %d, Type: %T\n", intValue, intValue)

	// (3) number -> string

	num := 42

	// Standard base-10 conversion
	str1 := strconv.Itoa(num)

	// Convert int64 to a base-10 string
	str2 := strconv.FormatInt(int64(num), 10)

	// Formatting inside another string
	str3 := fmt.Sprintf("The answer is %d", num)

	fmt.Println(str1, str2, str3)
}
