package main

import (
	"fmt"
)

func main() {

	k := make(chan bool, 2)

	go func() {
		k <- true
		k <- false
	}()

	fmt.Println(<-k, <-k)
	//output: true false
}
