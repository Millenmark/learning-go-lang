package main

import (
	"fmt"
	"math"
)

func main() {
	inflationRate := getUserInput("Enter an inflation rate: ")
	investmentAmount := getUserInput("Enter an investment amount: ")
	years := getUserInput("Enter a number of years: ")
	expectedReturnRate := getUserInput("Enter an expected return rate: ")

	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	futureValue, futureRealValue := calcValues(investmentAmount, expectedReturnRate, years, inflationRate)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for Inflation): %.2f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %.2f\nFuture Value (Adjusted for Inflation): %.2f", futureValue, futureRealValue)
	// fmt.Println("Future Value (Adjusted for Inflation): ", futureRealValue)

	// profit()

	
}

func calcValues(investmentAmount, expectedReturnRate, years, inflationRate float64) ( fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return
}

func getUserInput(text string) (userInput float64) {
	fmt.Print(text)
	fmt.Scan(&userInput)
	return
}