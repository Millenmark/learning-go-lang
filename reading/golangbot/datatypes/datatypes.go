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

	// Complex Data Type

	f := 5 + 3i
	fmt.Println("This is the value of f:", f)
	fmt.Printf("The type of f is %T\n", f)

	c1 := complex(5, 67)
	fmt.Println("The value of c1 is: ", c1)
	fmt.Printf("type of c1 is %T, size of c1 is %d bytes\n", c1, unsafe.Sizeof(c1))

	// constant
	// const ab = math.Sqrt(4) !this will throw an error, const cannot be assigned a value after declaration
	// fmt.Println(ab)

}
