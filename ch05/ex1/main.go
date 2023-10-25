package main

import (
	"errors"
	"fmt"
)

func main() {
	sumValue, _ := add(1, 2)
	fmt.Println("sum", sumValue)

	subValue, _ := sub(1, 2)
	fmt.Println("sub", subValue)

	divValue, _ := div(4, 2)
	fmt.Println("div", divValue)

	_, e := div(1, 2)
	fmt.Println("by 0", e)
}

func add(a int, b int) (int, error) {
	return a + b, nil
}

func sub(a int, b int) (int, error) {
	return a - b, nil
}

func mlt(a int, b int) (int, error) {
	return a * b, nil
}

func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
