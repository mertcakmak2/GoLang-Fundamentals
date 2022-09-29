package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	name := "mert cakmak"
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))             // string
	fmt.Printf("%d bytes\n", unsafe.Sizeof(name)) // 16 bytes

}
