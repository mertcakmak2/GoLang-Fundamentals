package main

import (
	"fmt"
	"time"
)

func main() {
	for range time.Tick(time.Second * 3) {
		time := time.Now().Format("15:04:05")
		fmt.Println("Time: " + time)
	}
}
