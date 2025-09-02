package main

import "fmt"

// Las constantes se declaran con la palabra reservada const y deben tener un valor inicial
// Ademas de venir en mayusculas

func main() {
	// Declaración de constantes
	const PI float64 = 3.14
	const NAME string = "Andres"

	// Declaración de constantes con tipo
	const (
		X float64 = 3.14
		Y float64 = 2.14
		Z float64 = 1.14
	)

	// Declaración de constantes con tipo y valor
	const (
		PI_3   = 3.14
		NAME_3 = "Andres"
	)

	const (
		MONDAY = iota + 1
		TUESDAY
		WEDNESDAY
		THURSDAY
		FRIDAY
		SATURDAY
		SUNDAY
	)

	fmt.Println(PI, NAME)
	fmt.Printf("X: %.2f, Y: %.2f, Z: %.2f\n", X, Y, Z)
	fmt.Println(PI_3, NAME_3)
	fmt.Println(MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY)
}
