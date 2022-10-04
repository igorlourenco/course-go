package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 13
	fmt.Println(s)

	s = make([]int, 10, 20)

	s = append(s, 1, 2, 3, 4, 5, 2, 1, 3, 54, 6)

	fmt.Println(s, len(s), cap(s))
}
