package main

import "fmt"

func main() {
	array1 := [3]int{3, 4, 5}
	var slice1 []int

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 4)
	copy(slice2, slice1) // first parameter receive the value copied from second parameter
	fmt.Println(slice2)
}
