package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d bytes\n", a, unsafe.Sizeof(a))
	fmt.Printf("type of b is %T, size of b is %d bytes\n", b, unsafe.Sizeof(b))

	var c uint = 56
	var d uint = 78
	e := c * d
	fmt.Printf("type of e is %T, size of e is %d bytes\n", e, unsafe.Sizeof(e))

}
