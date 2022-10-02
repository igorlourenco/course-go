package main

import "math"

func main() {
	a := 5
	b := 4

	println("+", a+b)
	println("-", a-b)
	println("x", a*b)
	println("/", float64(a)/float64(b))
	println("%", a%b)

	// bitwise operators
	println("&", a&b)
	println("|", a|b)
	println("^", a^b) // xor

	c := float64(2.0)
	d := float64(3.0)

	println(" > ", math.Max(c, d))
	println(" < ", math.Min(c, d))
	println(" ^ ", math.Pow(c, d))
}
