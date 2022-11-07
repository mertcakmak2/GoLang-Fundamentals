package main

import (
	"fmt"
	"sync"
	"time"
)

type Country struct {
	Name       string
	Continent  string
	Population int32
	mu         sync.Mutex
}

func updatePopulation(c *Country, newBorns int32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Population += newBorns
	fmt.Printf("New population of %v is %v\n", c.Name, c.Population)
}

func main() {
	totalStates := 20
	us := Country{"USA", "North America", 32000000, sync.Mutex{}}
	for i := 0; i < totalStates; i++ {
		go updatePopulation(&us, int32(1))
	}
	time.Sleep(time.Second * 5)
}
