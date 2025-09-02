package main

import (
	"errors"
	"fmt"
	"os"
)

func division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir por 0")
	}
	return a / b, nil
}

func main() {

	resultado, err := division(10, 2)
	if err != nil {
		fmt.Println("Error al dividir", err)
		return
	}
	fmt.Printf("El resultado de la división es: %d\n", resultado)

	// defer => se usa para ejecutar una función al final de la ejecución de la función actual
	file, err := os.Create("archivo.txt")

	if err != nil {
		fmt.Println("Error al crear el archivo", err)
		return
	}
	defer file.Close()

	file.WriteString("Hola mundo")

	// panic => se usa para lanzar un error y detener la ejecución del programa
	// panic("Error de prueba")

}
