package main

import "fmt"

func main() {
	var b byte
	var smallI int32
	var bigI uint64

	b = 255
	smallI = 2147483647
	bigI = 18446744073709551615

	b++
	smallI++
	bigI++

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
