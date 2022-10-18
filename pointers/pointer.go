package main

import "fmt"

func main() {

	user := User{Name: "mert", Lastname: "cakmak", Age: 24}
	// changeName(user)
	// fmt.Printf("user: %v\n", user)

	// output => user: {mert cakmak 24}

	changeNameWithPointer(&user)
	fmt.Printf("user: %v\n", user)
	// output => user: {changed name cakmak 24}
}

func changeName(user User) {
	user.Name = "changed name"
}

func changeNameWithPointer(user *User) {
	user.Name = "changed name"
}

type User struct {
	Name     string
	Lastname string
	Age      int
}
