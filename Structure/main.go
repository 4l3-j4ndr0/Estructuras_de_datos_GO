package main

import (
	"fmt"
)

// las estructuras funcionan como una clase que contiene atributos, 
	// a los cuales para acceder se debe instaciar primero dicha estructura

	// declaracion de estructura
	type Persona struct {
		nombre string
		edad   int
		correo string
	}

func main() {

	// asignacion de valores para los atributos de una estructura
	p := Persona{"Alex", 29, "alex@gmail.com"}
	fmt.Println(p)

	// edicion de valores de los atributos de una estructura
	p.edad = 30
	fmt.Println(p)

	// se puede realizar varias instancias de la misma estructura 
	p2 := Persona{"Adrian", 24, "adrian@gmail.com"}
	fmt.Println(p,p2)
}