package main

import "fmt"

func main() {
	var grades [3]float64
	grades[0], grades[1], grades[2] = 7.8, 4.3, 9.1

	fmt.Println(grades)

	var total float64 = 0.0
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	var average = total / float64(len(grades))

	fmt.Printf("Average: %.2f)", average)
}
