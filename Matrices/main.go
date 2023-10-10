package main

import (
	"fmt"
)

func main() {
	// matriz en blanco
	var matriz [5]int

	// matriz inicializada
	var matrizInicializada = [5]int{10, 20, 30, 40, 50}

	// matriz inicializada sin saber cuantos elementos seran se usa los 3 puntos
	var matrizInicializadaUnknow = [...]int{10, 20, 30, 40, 50}

	fmt.Println(matriz)
	fmt.Println(matrizInicializada)
	fmt.Println(matrizInicializadaUnknow)

	// 1 forma de iterar una matriz
	for i := 0; i < len(matrizInicializadaUnknow); i++ {
		fmt.Println(matrizInicializadaUnknow[i])
	}

	// forma de iterar una matriz con range
	for index, value := range matrizInicializadaUnknow {
		fmt.Printf("indice : %d , valor %d \n ", index, value)
	}

	// forma de iterar una matriz con range pero solo mostrando
	// los valores
	for _, value := range matrizInicializadaUnknow {
		fmt.Printf(" valor %d \n ", value)
	}

	// #############  MATRIZ BIDIMENSIONAL  #############

	// declarar
	var matrizBi = [3][3]int{}

	// iniializar
	var matrizBiInicializada = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	// puedo usar los 3 puntos pero solo en la fila , 
	// las columnas si deben estar definidas 	
	var matrizBiInicializadaUnknow = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	fmt.Println(matrizBi)
	fmt.Println(matrizBiInicializada)
	fmt.Println(matrizBiInicializadaUnknow)
}
