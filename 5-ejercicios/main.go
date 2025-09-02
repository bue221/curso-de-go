package main

import (
	"fmt"
	"math"
)

func calcularHipotenusa(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func calcularArea(a, b float64) float64 {
	return (a * b) / 2
}

func calcularPerimetro(a, b, c float64) float64 {
	return a + b + c
}

func imprimirResultado(area, perimetro, hipotenusa float64) {
	fmt.Printf("\nEl área del triángulo es: %.0f cm\n", area)
	fmt.Printf("El perímetro del triángulo es: %.0f cm\n", perimetro)
	fmt.Printf("La hipotenusa del triángulo es: %.0f cm\n", hipotenusa)
}

func main() {
	var a, b float64

	fmt.Println("Ingrese el lado a y b del triangulo en cm: ")
	fmt.Scanln(&a, &b)

	c := calcularHipotenusa(a, b)
	area := calcularArea(a, b)
	perimetro := calcularPerimetro(a, b, c)

	imprimirResultado(area, perimetro, c)

}
