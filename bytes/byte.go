package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var a1 byte = 97
	var a2 byte = 98
	var a3 byte = 99

	fmt.Println(a1, a2, a3)
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a1)) // 1 bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a2)) // 1 bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a3)) // 1 bytes

	fmt.Println("-------------------------")

	var a byte = 97
	var b = 98
	c := 'c'

	fmt.Println(a, b, c) // 97, 98, 99

	fmt.Println("-------------------------")

	fmt.Println(reflect.TypeOf(a)) // uint8
	fmt.Println(reflect.TypeOf(b)) // int
	fmt.Println(reflect.TypeOf(c)) // int32

}
