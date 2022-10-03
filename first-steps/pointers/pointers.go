package main

func main() {
	i := 1
	var p *int = nil

	p = &i // get the memory address of i

	println("i =", i, "p =", p, "*p =", *p)

	// i = 1 p = 0xc00004ff68 *p = 1
}
