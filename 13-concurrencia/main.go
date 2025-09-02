package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{"https://api.github.com", "https://api.google.com", "https://api.facebook.com"}

	for _, api := range apis {
		checkAPI(api)
	}

	elapsed := time.Since(start)

	fmt.Printf("Tiempo de ejecución: %s\n", elapsed)
}

func checkAPI(api string) {

	if _, err := http.Get(api); err != nil {
		fmt.Printf("Error al llamar a la API %s \n", api)
		return
	}

	fmt.Printf("La API %s esta funcionando \n", api)
}
