package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	name := []byte("mert")
	lastName := []byte("cakmak")

	fmt.Println(name, lastName)

	fmt.Printf("%d bytes\n", unsafe.Sizeof(name))     // 24 bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(lastName)) // 24 bytes

	fmt.Println(reflect.TypeOf(name))     // []uint8
	fmt.Println(reflect.TypeOf(lastName)) // []uint8

	fmt.Println("-------------------------")

	// byte to string
	n := string([]byte{109, 101, 114, 116})
	fmt.Println(n)
	fmt.Println(reflect.TypeOf(n))                // string
	fmt.Printf("%d bytes\n", unsafe.Sizeof(name)) // 24 bytes

}
