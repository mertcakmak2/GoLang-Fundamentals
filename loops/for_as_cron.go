package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("os.Args[1]: %v\n", os.Args[1])
	for range time.Tick(time.Second * 3) {
		time := time.Now().Format("15:04:05")
		fmt.Println("Time: " + time)
	}
}
