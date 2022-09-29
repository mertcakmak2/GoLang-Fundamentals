package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Signed Integers

	// int8 (8-bit signed integer whose range is -128 to 127)
	// int16 (16-bit signed integer whose range is -32768 to 32767)
	// int32 (32-bit signed integer whose range is -2147483648 to 2147483647)
	// int64 (64-bit signed integer whose range is -9223372036854775808 to 9223372036854775807)

	var int8num int8 = 127
	var int16num int16 = 32767
	var int32num int32 = 2147483647
	var int64num int64 = 9223372036854775807

	fmt.Println(int8num, int16num, int32num, int64num)
	fmt.Printf("%d bytes\n", unsafe.Sizeof(int8num))  // 1 bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(int16num)) // 2 bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(int32num)) // 4 bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(int64num)) // 8 bytes

	var int64num2 int64 = 1
	fmt.Printf("%d bytes\n", unsafe.Sizeof(int64num2)) // 8 bytes

	var a int = 1
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a)) // 8 bytes
}
