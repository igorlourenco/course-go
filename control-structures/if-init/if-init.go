package main

import (
	"math/rand"
	"time"
)

func getRandom() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)

}

func main() {
	if i := getRandom(); i > 5 {
		println("You won!")
	} else {
		println("You lost!")
	}
}
