package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(val int) {
			fmt.Println(val)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
