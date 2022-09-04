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
	fmt.Printf("%d bytes\n", unsafe.Sizeof(int8num))
	fmt.Printf("%d bytes\n", unsafe.Sizeof(int16num))
	fmt.Printf("%d bytes\n", unsafe.Sizeof(int32num))
	fmt.Printf("%d bytes\n", unsafe.Sizeof(int64num))

	var int64num2 int64 = 1
	fmt.Printf("%d bytes\n", unsafe.Sizeof(int64num2)) // 8 bytes

	// Unsigned integers in Go

	// uint8 (8-bit unsigned integer whose range is 0 to 255 )
	// uint16 (16-bit unsigned integer whose range is 0 to 65535 )
	// uint32 (32-bit unsigned integer whose range is 0 to 4294967295 )
	// uint64 (64-bit unsigned integer whose range is 0 to 18446744073709551615 )

	var uint8num uint8 = 255
	var uint16num uint16 = 65535
	var uint32num uint32 = 4294967295
	var uint64num uint64 = 18446744073709551615

	fmt.Println(uint8num, uint16num, uint32num, uint64num)
	fmt.Printf("%d bytes\n", unsafe.Sizeof(uint8num))
	fmt.Printf("%d bytes\n", unsafe.Sizeof(uint16num))
	fmt.Printf("%d bytes\n", unsafe.Sizeof(uint32num))
	fmt.Printf("%d bytes\n", unsafe.Sizeof(uint64num))

}
