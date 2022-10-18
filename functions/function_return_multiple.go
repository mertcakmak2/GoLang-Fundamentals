package main

import "fmt"

func main() {
	first, second := helloWorld()
	// first, _ := helloWorld() 	// Blank identifier (_)
	// _, second := helloWorld()	// Blank identifier (_)
	fmt.Println(first)
	fmt.Println(second)
}

func helloWorld() (string, string) {
	return "first hello world", "second hello world"
}
