package main

import "fmt"

func main() {

	user := User{Name: "mert", Lastname: "cakmak", Age: 24}

	user.printName()
	user.printAge()

}

type User struct {
	Name     string
	Lastname string
	Age      int
}

func (u User) printName() {
	fmt.Println(u.Name)
}

func (u User) printAge() {
	fmt.Println(u.Age)
}
