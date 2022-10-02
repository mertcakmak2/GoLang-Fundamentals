package main

import (
	"fmt"
	"time"
)

func main() {

	users := []string{"user 1", "user 2", "user 3", "user 4"}
	channel := make(chan string)

	fmt.Println("before")
	go sender(channel, users)
	go receiver(channel, users)
	fmt.Println("after")

	<-time.After(time.Second * 5)
}

func sender(channel chan string, users []string) {
	func(userList []string) {
		for _, user := range userList {
			fmt.Println("Sent user: " + user)
			channel <- user // send to channel
			time.Sleep(time.Second)
		}
	}(users)
}

func receiver(channel chan string, users []string) {
	func(userList []string) {
		for i := 0; i < len(userList); i++ {
			user := <-channel
			fmt.Println("Received user: " + user)
		}
	}(users)
}

// Output

// before
// after
// Sent user: user 1
// Received user: user 1
// Sent user: user 2
// Received user: user 2
// Sent user: user 3
// Received user: user 3
// Sent user: user 4
// Received user: user 4