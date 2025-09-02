package main

import (
	"fmt"
)

func main() {

	firstName, lastName, age, isStudent := "Andres", "Plaza", 20, false

	fmt.Println(firstName, lastName, age, isStudent)

	firstName = "Juan"
	lastName = "Perez"
	age = 30
	isStudent = true

	fmt.Println(firstName, lastName, age, isStudent)
}
