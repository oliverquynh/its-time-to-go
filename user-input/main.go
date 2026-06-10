package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	for {
		fmt.Print("Please enter an integer: ")

		// Pass the memory address using & so the function can modify the variable
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("\n  [ERROR] Could not read your input:", err)
            fmt.Println("")
			return
		}

		n, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("\n  [ERROR] Your input is not a valid integer:", err)
            fmt.Println("")
            continue
		}

		fmt.Printf("\n  > Your integer is %d.\n\n", n)
	}
}
