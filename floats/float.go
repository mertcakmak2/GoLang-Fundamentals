package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var f float32
	var f2 float64
	f = 12.34567890123456
	f2 = 12.34567890123456

	fmt.Println(f, f2)              // "12.345679 12.34567890123456"
	fmt.Println(reflect.TypeOf(f))  // float32
	fmt.Println(reflect.TypeOf(f2)) // float64

	fmt.Printf("%d bytes\n", unsafe.Sizeof(f))  // 4 bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(f2)) // 8 bytes
}
