package main

import "fmt"

func getGradeConcept(grade float64) {

	_grade := int(grade)

	switch _grade {
	case 10, 9:
		fmt.Println("A")
	case 8, 7:
		fmt.Println("B")
	case 6, 5:
		fmt.Println("C")
	case 4, 3:
		fmt.Println("D")
	case 2, 1:
		fmt.Println("E")
	case 0:
		fmt.Println("F")
	default:
		fmt.Println("Invalid grade")
	}
}

func main() {
	getGradeConcept(9.8)
	getGradeConcept(3)
	getGradeConcept(7)

}
