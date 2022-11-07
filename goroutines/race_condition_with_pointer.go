package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println(time.Now())

	wg := sync.WaitGroup{}
	wg.Add(2)
	lock := sync.Mutex{}
	val := 0
	go firstIncrementMethod(&val, &wg, &lock)
	go secondIncrementMethod(&val, &wg, &lock)
	wg.Wait()

	fmt.Println(val)
	fmt.Println(time.Now())
}

func firstIncrementMethod(val *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	time.Sleep(5 * time.Second)
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*val++
		lock.Unlock()
	}
	wg.Done()
}

func secondIncrementMethod(val *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	time.Sleep(5 * time.Second)
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*val++
		lock.Unlock()
	}
	wg.Done()
}
