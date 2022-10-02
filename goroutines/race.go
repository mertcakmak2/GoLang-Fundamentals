package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	lock := sync.Mutex{}
	val := 0

	go func() {
		for i := 0; i < 100000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(val)
}
