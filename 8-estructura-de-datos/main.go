package main

import "fmt"

func main() {
	// Arrays
	var array [5]int
	matriz := [...]int{1, 2, 3}

	fmt.Println(array, matriz)

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	fmt.Println(array)

	for index, value := range matriz {
		fmt.Println(index, value)
	}

	// Slices
	diasSemana := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}

	entreSemana := diasSemana[0:5]
	finesemana := diasSemana[5:7]

	fmt.Println(entreSemana, finesemana)

	for index, value := range diasSemana {
		fmt.Println(index, value)
	}

	// Maps
	paises := make(map[string]string)

	paises["ES"] = "España"
	paises["FR"] = "Francia"
	paises["DE"] = "Alemania"
	paises["IT"] = "Italia"
	paises["PT"] = "Portugal"

	fmt.Println(paises)
}
