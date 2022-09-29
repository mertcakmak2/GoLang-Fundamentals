package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("start")

	sumResult := make(chan int, 1)
	divideResult := make(chan int, 1)

	// sumValue := sum(2, 3, sumResult)
	// dividedValue := divide(12, 3, divideResult)
	// fmt.Printf("sumValue: %d\n", sumValue)
	// fmt.Printf("dividedValue: %d\n", dividedValue)

	// completed in 4 second

	go sum(2, 3, sumResult)
	go divide(12, 3, divideResult)
	sumValue := <-sumResult
	dividedValue := <-divideResult
	fmt.Printf("sumValue: %d\n", sumValue)
	fmt.Printf("dividedValue: %d\n", dividedValue)

	// completed in 2 second

	<-time.After(time.Second * 5)
}

func sum(a, b int, result chan int) int {
	time.Sleep(time.Second * 2)
	sumValue := a + b
	result <- sumValue
	return sumValue
}

func divide(a, b int, result chan int) int {
	time.Sleep(time.Second * 2)
	dividedValue := a / b
	result <- dividedValue
	return dividedValue
}
