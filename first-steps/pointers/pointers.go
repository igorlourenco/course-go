package main

func main() {
	i := 1

	// go has no pointer arithmetic, so you can't do i++ or i--

	var p *int = nil

	p = &i // get the memory address of i

	println("i =", i, "p =", p, "*p =", *p)
}
