package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5
	// var years float64 = 10

	fmt.Print("Enter your investment amount: ")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future value: %v", futureRealValue)

	fmt.Println("Future value (adjusted for inflation): ", futureRealValue)
	outputText("futureRealValue")
}

func outputText(text string) {
	fmt.Print("Hellpopppp", text)
}
