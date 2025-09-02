package main

import (
	"fmt"
	"os"
)

type Tarea struct {
	id          int
	nombre      string
	descripcion string
	completada  bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (l *ListaTareas) agregarTarea(tarea Tarea) {
	l.tareas = append(l.tareas, tarea)
}

func (l *ListaTareas) marcarCompletada(id int) {
	for i := range l.tareas {
		if l.tareas[i].id == id {
			l.tareas[i].completada = true
			break
		}
	}
}

func (l *ListaTareas) eliminarTarea(id int) {
	for i, tarea := range l.tareas {
		if tarea.id == id {
			l.tareas = append(l.tareas[:i], l.tareas[i+1:]...)
			break
		}
	}
}

func (l *ListaTareas) verTareas() {
	fmt.Println("\n------------TAREAS---------------")
	if len(l.tareas) == 0 {
		fmt.Println("No hay tareas")
	} else {
		for _, tarea := range l.tareas {
			fmt.Printf("ID: %d, Nombre: %s, Descripción: %s, Completada: %t\n", tarea.id, tarea.nombre, tarea.descripcion, tarea.completada)
			fmt.Println("--------------------------------\n	")
		}
	}
}

func printMenu() (int, error) {
	fmt.Println("\n------------MENU---------------")
	fmt.Println("1. Agregar tarea")
	fmt.Println("2. Marcar tarea como completada")
	fmt.Println("3. Eliminar tarea")
	fmt.Println("4. Ver todas las tareas")
	fmt.Println("5. Salir")
	fmt.Println("Ingrese un número: ")
	var response int
	_, err := fmt.Scanln(&response)
	fmt.Println("------------FIN MENU---------------")
	return response, err
}

func main() {
	listaTareas := ListaTareas{}

	for {
		response, err := printMenu()
		if err != nil {
			fmt.Println("Opción inválida")
			continue
		}

		switch response {
		case 1:
			fmt.Println("Ingrese el nombre, descripcion de la tarea separados por espacio: ")
			var nombre, descripcion string
			fmt.Scanln(&nombre, &descripcion)
			listaTareas.agregarTarea(Tarea{nombre: nombre, descripcion: descripcion, completada: false, id: len(listaTareas.tareas) + 1})
		case 2:
			fmt.Println("Ingrese el ID de la tarea a marcar como completada: ")
			var id int
			fmt.Scanln(&id)
			listaTareas.marcarCompletada(id)
		case 3:
			fmt.Println("Ingrese el ID de la tarea a eliminar: ")
			var id int
			fmt.Scanln(&id)
			listaTareas.eliminarTarea(id)
			fmt.Println("Tarea eliminada")
		case 4:
			listaTareas.verTareas()
		case 5:
			fmt.Println("Saliendo...")
			os.Exit(0)
		default:
			fmt.Println("Opción inválida")
		}
	}

}
