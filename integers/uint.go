package main

import (
	"fmt"
	"unsafe"
)

func main() {
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
	fmt.Printf("%d bytes\n", unsafe.Sizeof(uint8num))  // 1 bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(uint16num)) // 2 bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(uint32num)) // 4 bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(uint64num)) // 8 bytes

}
