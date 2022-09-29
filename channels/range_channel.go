package main

import (
	"fmt"
)

func main() {

	ch := make(chan string, 3)

	ch <- "mert"
	ch <- "cakmak"
	ch <- "trabzon"
	close(ch)

	for word := range ch {
		fmt.Println(word)
	}

	// Output

	// mert
	// cakmak
	// trabzon
}
