package main

import "fmt"

func main() {
	taxRate := 0.2

	result := calculateBill(23, 2)
	fmt.Println(result)

	// closures (anonymous function)
	calculateTax := func(price float64) float64 {
		return price * taxRate
	}

	tax := calculateTax(float64(result))
	fmt.Println(tax)

	rectProps := func(width, height float64) (float64, float64) {
		return width * height, width + height
	}

	area, perimeter := rectProps(4, 5)

	fmt.Println("the area is" + " " + fmt.Sprintf("%f", area) + " " + fmt.Sprintf("%f", perimeter))
}
