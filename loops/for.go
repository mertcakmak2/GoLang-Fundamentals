package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	numbers := [6]int{2, 3, 5, 7, 11, 13}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
