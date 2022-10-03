package main

import "time"

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("Good morning!")
	case t.Hour() < 18:
		println("Good afternoon!")
	default:
		println("Good evening!")
	}
}
