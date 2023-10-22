package main

import "fmt"

func main() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	var greetings1 = greetings[:2]
	var greetings2 = greetings[1:4]
	var greetings3 = greetings[3:]

	fmt.Println("greetings1:", greetings1)
	fmt.Println("greetings2:", greetings2)
	fmt.Println("greetings3:", greetings3)
}
