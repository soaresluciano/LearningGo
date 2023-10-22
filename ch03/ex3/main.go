package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	literal_without_keys := Employee{
		"Tom",
		"Bombadil",
		1,
	}

	literal_with_keys := Employee{
		firstName: "Frodo",
		lastName:  "Baggins",
		id:        2,
	}

	var dot_notation Employee

	dot_notation.firstName = "Pippin"
	dot_notation.lastName = "Took"
	dot_notation.id = 3

	fmt.Println(literal_without_keys)
	fmt.Println(literal_with_keys)
	fmt.Println(dot_notation)
}
