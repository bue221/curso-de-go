package main

import "fmt"

// Variadic functions
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T: %v\n", dato, dato)
	}
}

// Recursive functions
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Higher order functions
func sumarWrapper(a, b int, fn func(int ...int) int) int {
	return fn(a, b)
}

// Closures
func incrementar() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	imprimirDatos("Andres", 20, true)
	imprimirDatos(1, 2, 3, 4, 5)
	fmt.Println(factorial(5))

	func() {
		fmt.Println("Soy una función anónima")
	}()

	fmt.Println(sumarWrapper(1, 2, sum))

	incrementar := incrementar()
	fmt.Println(incrementar())
	fmt.Println(incrementar())
}
