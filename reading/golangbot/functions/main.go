package main

import "fmt"

func main() {
	taxRate := 0.2

	result := calculateBill(23, 2)
	fmt.Println(result)

	// closures
	calculateTax := func(price float64) float64 {
		return price * taxRate
	}

	tax := calculateTax(float64(result))
	fmt.Println(tax)
}
