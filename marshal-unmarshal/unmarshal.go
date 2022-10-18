package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user_json := `{
		"name" : "mert",
		"lastName" : "cakmak",
		"age" : 24
	}`
	var user User
	bytes := []byte(user_json)
	json.Unmarshal(bytes, &user)
	fmt.Println(user) // {mert cakmak 24}
}

type User struct {
	Name     string `json:"name"`
	Lastname string `json:"lastName"`
	Age      int    `json:"age"`
}
