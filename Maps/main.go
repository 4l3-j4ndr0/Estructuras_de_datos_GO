package main

import (
	"fmt"
)

func main() {
	// declara e inicializar un map
	color := map[string]string {
		"rojo" : "#FF0000",
		"verde" : "#00FF00",
		"azul" : "#0000FF",
	}

	// editAR UN ELEMENTO
	color["rojo"] = "sangre"

	// agregar un elemento 
	color["blanco"] = "#FFFFFF"
	fmt.Println(color) 

	//tomar vlor de un map
	value:= color["rojo"]
	fmt.Println(value)

	// tomar valor de un map con un verificador
	// si existe devolvera el valor y true , si no existe 
	// devuelve un string vacio y false 
	value2, ale:= color["rojo"]
	fmt.Println(value2,ale)

	// eliminar un elemento de un map
	delete(color,"rojo")
	fmt.Println(color)

	// iterar con for
	for clave,valor := range color {
		fmt.Printf("clave : %s , valor : %s \n",clave,valor)
	}
}