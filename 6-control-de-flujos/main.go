package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	hour := t.Hour()

	if t.Weekday() == time.Monday {
		fmt.Println("Hoy es lunes")
	} else {
		fmt.Println("Hoy no es lunes")
		if hour > 12 {
			fmt.Println("Es tarde")
		} else {
			fmt.Println("Es temprano")
		}
	}

	switch t.Weekday() {
	case time.Monday:
		fmt.Println("Hoy es lunes")
	case time.Tuesday:
		fmt.Println("Hoy es martes")
	case time.Wednesday:
		fmt.Println("Hoy es miércoles")
	default:
		fmt.Println("No es tu día de trabajo")
	}

}
