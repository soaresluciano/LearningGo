package main

import (
	"fmt"
	"math/rand"
)

func main() {
	max := 100
	s := make([]int, max)

	for i := 0; i < max; i++ {
		value := rand.Intn(max)
		s[i] = value

		PrintMessage(i, value)
	}
}

func PrintMessage(i int, v int) {
	divBy2 := v%2 == 0
	divBy3 := v%3 == 0
	both := divBy2 && divBy3

	var message string

	switch {
	case both:
		message = "Six!"
	case divBy3:
		message = "Three!"
	case divBy2:
		message = "Two!"
	default:
		message = "Never mind"
	}

	fmt.Println(i, v, message)
}
