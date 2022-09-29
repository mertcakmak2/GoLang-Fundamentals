package main

import "fmt"

func main() {

	numbers := [6]int{2, 3, 5, 7, 11, 13}

	for _, n := range numbers {
		fmt.Println(n)
	}
}
