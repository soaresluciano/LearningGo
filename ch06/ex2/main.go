package main

import "fmt"

func main() {
	target := make([]string, 5)

	PrintSlice(target)
	UpdateSlice(target, "Luciano")
	PrintSlice(target)

	PrintSlice(target)
	newArray := GrowSlice(target, "Lulu")
	PrintSlice(newArray)
}

func UpdateSlice(target []string, value string) {
	lastPosition := len(target) - 1
	target[lastPosition] = value
}

func GrowSlice(target []string, value string) []string {
	return append(target, value)
}

func PrintSlice(target []string) {
	for i := 0; i < len(target); i++ {
		ReportSlice(i, target[i])
	}
}

func ReportSlice(index int, value string) {
	fmt.Printf("%-20v\t%-2v\n", index, value)
}

func Report(desc string, value string) {
	fmt.Printf("%-20v\t%-2v\n", desc, value)
}
