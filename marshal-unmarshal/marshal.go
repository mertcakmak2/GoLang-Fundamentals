package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	user := User{Name: "mert", Lastname: "cakmak", Age: 24}
	u, _ := json.Marshal(user)

	fmt.Println(user)      // {mert cakmak 24}
	fmt.Println(string(u)) // {"name":"mert","lastName":"cakmak","age":24}
}

type User struct {
	Name     string `json:"name"`
	Lastname string `json:"lastName"`
	Age      int    `json:"age"`
}
