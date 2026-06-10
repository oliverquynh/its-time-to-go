package main

import "fmt"

func main() {
	values := [7]int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(values); i++ {
		fmt.Printf("%d\n", values[i])
	}

	fmt.Println(values[1:3]) // [1 2]

	weekDays := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for _, day := range weekDays {
		fmt.Println(day)
	}
}
