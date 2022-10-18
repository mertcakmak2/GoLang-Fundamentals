package main

import (
	"fmt"
	"reflect"
)

func main() {

	arr := [2]string{"I", "learn"}
	fmt.Printf("arr: %v\n", arr)
	fmt.Println(reflect.TypeOf(arr)) // [2]string

	l := len(arr)
	a := append(arr[:l], "golang")
	fmt.Printf("a: %v\n", a)

	fmt.Println(reflect.TypeOf(a)) // []string

}
