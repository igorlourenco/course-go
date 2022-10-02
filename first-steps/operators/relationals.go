package main

func main() {
	println(" == ", "banana" == "banana")
	println(" != ", "banana" != "banana")
	println("<", 3 < 2)
	println(">", 3 > 2)
	println("<=", 3 <= 2)
	println(">=", 3 >= 2)

	type Person struct {
		Name string
		Age  uint8
	}

	p1 := Person{"John", 20}
	p2 := Person{"John", 20}

	println("person == person", p1 == p2)
}
