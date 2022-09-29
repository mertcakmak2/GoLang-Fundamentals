package main

import (
	"fmt"
	"strings"
)

func main() {

	w1 := "mert"
	w2 := "mert"

	if w1 == w2 {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
	// equal

	if strings.EqualFold(w1, w2) {
		fmt.Println("The words are equal")
	} else {
		fmt.Println("The words are not equal")
	}
	// The words are equal
}
