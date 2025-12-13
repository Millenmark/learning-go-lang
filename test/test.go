package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c int

	// array
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	const height int = 34
	// apple, mango := 12, "new"

	// mango = 89
	fmt.Println(height)

	const (
		DECIMAL     = 255
		OCTAL       = 0377
		HEXADECIMAL = 0xFF
	)

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(1 == '1')

	/* Bitwise Operators */
	C := 30

	// Bitwise AND assignment
	C &= 2
	fmt.Println("C &= 2 =", C)

	// Bitwise XOR assignment
	C ^= 2
	fmt.Println("C ^= 2 =", C)

	// Bitwise OR assignment
	C |= 2
	fmt.Println("C |= 2 =", C)

	// Bitwise AND NOT assignment
	C &^= 2
	fmt.Println("C &^= 2 =", C)

	// AND
	fmt.Println(1 & 2)

	// OR
	fmt.Println(1 | 2)

	// XOR
	fmt.Println(1 ^ 2)

	// NOT
	fmt.Println(^1)

	// AND NOT
	fmt.Println(1 &^ 2)

	// Left Shift
	fmt.Println(1 << 2)

	// Right Shift
	fmt.Println(1 >> 2)

	/* Miscellaneous Operators */
	A := 10

	// "Address of" operator
	ptr := &A
	// memory address of A
	fmt.Println("Address of A: ", ptr)

	//Pointer Dereference operator
	fmt.Println("Value of *ptr: ", *ptr)

}

/* Control Flow */
func loops() {
	fmt.Println("=== 1. Classic for loop ===")
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	fmt.Println("\n=== 2. Nested loop with labels ===")

Outer:
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			if x == 2 && y == 2 {
				fmt.Println("  break Outer at (2,2)")
				break Outer // jump out of BOTH loops
			}
			fmt.Printf("  (%d,%d)\n", x, y)
		}
	}

	fmt.Println("\n=== 3. While-style loop (condition only) ===")
	n := 1
	for n < 5 { // no parens, no semicolons
		fmt.Println("n =", n)
		n++
	}

	fmt.Println("\n=== 4. Do-while-style loop (run at least once) ===")
	k := 0
	for {
		fmt.Println("k =", k)
		k++
		if k >= 3 {
			break // test after the body
		}
	}

	fmt.Println("\n=== 5. Infinite loop with break/continue ===")
	sum := 0
	for {
		sum++
		if sum%7 == 0 {
			continue // skip the rest of this iteration
		}
		if sum > 15 {
			break // leave the loop
		}
		fmt.Println("sum =", sum)
	}

	fmt.Println("\n=== 6. Goto (rarely used, but valid) ===")
	counter := 0
Start:
	fmt.Println("counter =", counter)
	counter++
	if counter < 3 {
		goto Start // jump back up
	}
}

// ---------- classic recursion: factorial ----------
func fact(n int) int {
	if n <= 1 { // base case
		return 1
	}
	return n * fact(n-1) // recursive call
}

// ---------- another classic: Fibonacci ----------
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

/* NOTE: You can declare another function inside of a function but it should be anonymous */
func outer() {
	add := func(a, b int) int { return a + b }
	fmt.Println(add(3, 4))
}

var copy, pish string

// immediate call function
func outer1() {
	func(msg string) {
		fmt.Println(msg)
	}("hello from inside")
}

// closure
func makeAdder(x int) func(int) int {
	return func(y int) int { return x + y }
}

func learningArray() {
	// 1. Declare (zero-initialized)
	var balance [10]float32
	fmt.Println(balance) // [0 0 0 0 0 0 0 0 0 0]

	// 2. Declare + initialize in one line
	balance2 := [5]float32{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(balance2) // [1 2 3 4 5]

	// 3. Let compiler count for you
	balance3 := [...]float32{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(balance3) // [1 2 3 4 5]
}
