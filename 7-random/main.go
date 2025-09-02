package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := rand.Intn(100) + 1

	var input int

	for input != random {
		fmt.Println("Ingrese un número: ")
		fmt.Scanln(&input)

		if input == random {
			fmt.Println("¡Correcto!")
			break
		} else if input > random {
			fmt.Println("El número es menor")
		} else {
			fmt.Println("El número es mayor")
		}
	}
}
