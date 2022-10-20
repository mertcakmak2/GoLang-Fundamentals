package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	forever := make(chan bool)

	go func() { c1 <- "first channel message" }()
	go func() { c2 <- "second channel message" }()
	go func() { c2 <- "second channel message v2" }()
	go func() { c1 <- "first channel message v2" }()
	fmt.Printf("runtime.NumGoroutine(): %v\n", runtime.NumGoroutine()) // 5 (4 go func + main = 5)

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			select {
			case firstChannelMessage := <-c1:
				fmt.Println("received message from first channel", firstChannelMessage)
			case secondChannelMessage := <-c2:
				fmt.Println("received message from second channel", secondChannelMessage)
			default:
				// fmt.Println("no value received")
				// return
				forever <- true
			}
		}
	}()
	// time.Sleep(time.Second * 2)
	<-forever
}
