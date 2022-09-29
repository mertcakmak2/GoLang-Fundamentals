package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	content, err := ioutil.ReadFile("text.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(content)
	fmt.Println("-------------")
	fmt.Println(string(content))

	// 	Output

	// 	[109 101 114 116 13 10 99 97 107 109 97 107 13 10 116 114 97 98 122 111 110]
	// -------------
	// mert
	// cakmak
	// trabzon

}
