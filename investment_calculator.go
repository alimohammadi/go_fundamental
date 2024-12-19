package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5
	// var years float64 = 10

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	fmt.Print(futureValue)
}
