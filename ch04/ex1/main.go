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
		fmt.Println(i, value)
	}
}
