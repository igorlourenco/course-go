package main

import "time"

func main() {

	// use for like a while
	i := 0

	for i < 10 {
		println(i)
		i++
	}

	// use for like a for
	for j := 0; j < 10; j += 2 {
		println(j)
	}

	// infinite loop
	for {
		println("infinite")
		time.Sleep(time.Second * 2)
	}
}
