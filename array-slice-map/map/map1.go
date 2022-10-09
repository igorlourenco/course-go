package main

import "fmt"

func main() {
	// var approved map[int]string
	// maps must be initialized

	approved := make(map[int]string)
	approved[1] = "João"
	approved[4] = "Maria"
	approved[6] = "Ana"

	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Println(cpf, name)
	}

	fmt.Printf("approved 6: %s \n", approved[6])

	delete(approved, 6)
	fmt.Printf("approved 6: %s", approved[6])
}
