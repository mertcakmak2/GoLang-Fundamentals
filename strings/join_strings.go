package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	w1 := "learn"
	w2 := "golang"

	msg := fmt.Sprintf("%s %s", w1, w2)

	fmt.Println(msg) // learn golang

	var buf bytes.Buffer

	buf.WriteString("I am ")
	buf.WriteString("learning ")
	buf.WriteString("golang")

	fmt.Println(buf.String()) // I am learning golang

	words := []string{"I", "learn", "golang"}
	joinwords := strings.Join(words, " ")

	fmt.Println(joinwords) // I learn golang
}
