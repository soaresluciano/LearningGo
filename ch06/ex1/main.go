package main

func main() {
	MakePerson("John", "Doe", 32)
	MakePersonPointer("Ann", "Moe", 24)
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}
