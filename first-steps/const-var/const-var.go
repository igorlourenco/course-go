package main

import (
	"fmt"
	"math"
)

func main() {
	const pi float64 = 3.14159265359
	var radius = 3.2 // type float64

	area := pi * math.Pow(radius, 2)

	fmt.Println("Area =", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	var g, h, i, j int = 5, 6, 7, 8

	k, l := true, "variable"

	fmt.Println(a, b, c, d, g, h, i, j, k, l)
}
