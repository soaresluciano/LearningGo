package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	var runes []rune = []rune(message)
	var r4 rune = runes[3]
	var s4 string = string(r4)
	fmt.Println(s4)
}
