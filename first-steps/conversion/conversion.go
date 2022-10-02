package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	// fmt.Println(x / y) // error
	fmt.Println(x / float64(y)) // error

	// float to int
	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade) // 6

	// int to string? no
	num := 123
	fmt.Println("This is a number: " + string(num)) // error
	// this is not a number: 123, it's a correspondent character
	// in the ASCII table

	// int to string
	fmt.Println("This is a number: " + strconv.Itoa(123)) // This is a number: {

	// string to int
	num, _ = strconv.Atoi("123")
	// the second return is an error in case of
	// failure (sending a string) that's not a number, e.g. "123a"
	fmt.Println(num - 122) // 1

	// string to bool
	b, _ := strconv.ParseBool("true")
	// true or 1 -> true
	if b {
		fmt.Println("It's true!")
	}
}
