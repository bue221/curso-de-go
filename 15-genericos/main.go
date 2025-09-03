package main

import "fmt"

func PrintList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}

// Restricciones de tipo
type Numbers interface {
	~int | ~float64
}

func Sum[T Numbers](list ...T) T {
	var total T
	for _, item := range list {
		total += item
	}
	return total
}

func Includes[T comparable](list []T, item T) bool {
	for _, item := range list {
		if item == item {
			return true
		}
	}
	return false
}

func main() {
	PrintList(1, 2, 3, 4, 5)
	PrintList("Hello", "World")
	PrintList(true, false)
	fmt.Println(Sum(1, 2, 3, 4, 5.5))

	strings := []string{"Hello", "World", "Go"}
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Includes(strings, "Hello"))
	fmt.Println(Includes(numbers, 1))
}
