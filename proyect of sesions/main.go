package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

func (l *ListaTareas) editarTareas(index int, t Tarea) {
	l.tareas[index] = t
}

func (l *ListaTareas) eliminarTareas(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	lista := ListaTareas{}

	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("Seleccione una opcion  \n",
			"1- Agregar tarea  \n",
			"2- Marcar tarea como completada  \n",
			"3- Editar atarea  \n",
			"4- Eliminar tarea  \n",
			"5- Salir")
		fmt.Println("Ingrese la opcion")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Nombre de la tarea")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Descripcion de la tarea")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")

		case 2:
			var index int
			fmt.Println("Indice de la tarea a marcar completa")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente")

		case 3:
			var index int
			var t Tarea
			fmt.Println("Indice de la tarea a editar")
			fmt.Scanln(&index)
			fmt.Println("Nuevo nombre de la tarea")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Nueva descripcion de la tarea")
			t.descripcion, _ = leer.ReadString('\n')
			lista.editarTareas(index, t)
			fmt.Println("Tarea editada correctamente")

		case 4:
			var index int
			fmt.Println("Indice de la tarea a eliminar")
			fmt.Scanln(&index)
			lista.eliminarTareas(index)
			fmt.Println("Tarea eliminada correctamente")

		case 5:
			fmt.Println("Saliendo del programa ...")
			return

		default:
			fmt.Println("Opcion no valida")

		}
		fmt.Println("LISTA DE TAREAS")
		fmt.Println("==============================================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - completado %t \n", i, t.nombre, t.descripcion, t.completado)
		}
	}
}
