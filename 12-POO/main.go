package main

import (
	"library/animal"
)

func main() {
	// myBook := book.Book{
	// 	ID:     1,
	// 	Title:  "El principito",
	// 	Author: "Antoine de Saint-Exupéry",
	// 	Pages:  120,
	// }
	// myTextBook := book.NewTextBook("El principito 3", "Antoine de Saint-Exupéry", "El principito", "Editorial 3", "Secundaria", 3, 120)

	// book.Print(myTextBook)
	// book.Print(&myBook)

	miPerro := animal.Perro{Nombre: "Firulais"}
	miGato := animal.Gato{Nombre: "Michi"}

	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

}
