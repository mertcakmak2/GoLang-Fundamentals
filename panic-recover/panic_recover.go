package main

import (
	"fmt"
	"time"
)

func main() {

	printArrElement()
	fmt.Println("normally returned from main")

	// Output:

	// Recovered runtime error: index out of range [4] with length 3
	// normally returned from main

	time.Sleep(time.Hour)
}

func printArrElement() {
	defer recoverFunc()
	n := []int{5, 7, 4}
	fmt.Println(n[4])

}

func recoverFunc() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}
