package main

func main() {
	var b byte = 3
	println("b =", b)

	i := 3

	i += int(b)
	i -= 2
	i *= 3
	i /= 2

	println("i =", i)

	x, y := 1, 2

	println("x =", x, "y =", y)

	x, y = y, x

	println("x =", x, "y =", y)
}
