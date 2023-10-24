package main

import (
	"fmt"
)
type Persona struct {
	nombre string
	edad   int
	correo string
}
func main() {

	// ################  PUNTEROS  ################

	var x int = 10

	// para declarar unpuntero se usa la sintaxys : nombre_variable + *tipo_dato
	var pp *int = &x
	fmt.Println(pp)
	fmt.Println(&x)

	editar(&x)
	fmt.Println(&x)

	// para llamar a un metodo para q se ejecute no se llama como una funcion 
	// sino como si fuera un atributo mas de las estructura
	p := Persona{"Alex", 29,"alex@gmail.com"}
	p.diHola()
}

func editar(x *int) {
	*x = 20
}

// ##################  Metodos  ##################

// un metodo es una funcion pero se debe declarar un receptor antes del nombre de este metodo
// en este caso es la intancia p de la estructura Persona
func (p *Persona) diHola(){
	fmt.Println("Hola , mi nombre es : ",p.nombre)
}