package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{"https://api.github.com", "https://api.google.com", "https://api.facebook.com"}

	ch := make(chan string)
	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for range apis {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)

	fmt.Printf("Tiempo de ejecución: %s\n", elapsed)
}

func checkAPI(api string, ch chan string) {

	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("Error al llamar a la API %s \n", api)

		return
	}

	ch <- fmt.Sprintf("La API %s esta funcionando \n", api)
}
