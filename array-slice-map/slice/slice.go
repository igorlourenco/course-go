package main

import "fmt"

func main() {
	a1 := [5]int{2, 3, 4, 5, 6} // array
	// s1 := []int{4, 5, 6}        // slice

	s2 := a1[1:3]
	fmt.Println(s2)

}
