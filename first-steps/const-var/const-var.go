package main

import "fmt"

func main() {
	const pi float64 = 3.14159265359
	// var radius = 3.2 // type float64

	// area := pi * math.Pow(radius, 2)

	convertedPi := fmt.Sprint(pi)
	fmt.Println("Area = " + convertedPi)

}
