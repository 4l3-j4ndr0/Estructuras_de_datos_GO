package main

import (
	"fmt"
)

func main() {

	// declarar un slice 
	var slice [] string 
	fmt.Println("slice : ",slice)

	// declara e inicializar un slice
	diasSemana := [] string {"Domingo", "lunes", "Martes", "Miercoles", 
		"Jueves", "Viernes", "sabado"}
		fmt.Println("diasSemana : ",diasSemana)

	// crear un slice a partir de otro slice
	// dentro del corchete se establece el intervalo que se copiara del 
	// slice original, OJO : el primer valor (en este caso cero) si se toma en cuenta , 
	// pero el ultimo valor (en este caso 5) no se toma , se queda uno antes
	// Salida 
	// slice original : [Domingo lunes Martes Miercoles Jueves Viernes sabado]
	// slice copia : [Domingo lunes Martes Miercoles Jueves]
	// Nota : si queremos copiar integro todos los valores desde el original a la copia
	// la sintaxis es la misma pero dentro del corchete no ponemos el valor final
	// ejemplo : copyDiasSemana := diasSemana[0:]
	copyDiasSemana := diasSemana[0:5]
	fmt.Println("copyDiasSemana : ",copyDiasSemana)

	// agregar elementos a slice, sintax : 
	// A = append(A,B)
	// A --> donde se guardaran los cambios
	// A dentro de append --> es el slice que se tomara de base para agregar los demas valores
	// B --> valores a agregar. Pueden ser valores por separado (value1, value2, value3),
	// 1 valor de otro slice (slice[2]), otro slice completo (otherSlice...)  
	copyDiasSemana = append(copyDiasSemana, diasSemana[5])
	fmt.Println("copyDiasSemana agregado : ",copyDiasSemana)

	// para eliminar elementos de un slice se usa el mismo append
	// ejemplo : A = [1,2,3,4,5], se quiere eliminar 
	// la posicion 2 del slice A que seria el valor 3, seria de la siguiente forma :
	// A = append(A[:2],A[3:]...)
	// A --> donde se guardaran los cambios
	// A[:2] --> tomara de base el slice A desde el inicio hasta las posicion 2 despreciando esta posicion 2
	// A[3] --> agregarael resto del slice partiendo de la posicion 3 
	copyDiasSemana = append(diasSemana[:2],diasSemana[3:]...)
	fmt.Println("copyDiasSemana eliminando : ",copyDiasSemana)

	// Se puede crear un slice usando make, sintax : 
	// A := make([]string ,5 , 10)
	// []string --> tipo de dato del slice
	// 5 --> longitud
	// 10 --> capacidad
    // Longitud: La longitud es el número total de elementos presentes en el Array.
    // Capacidad: La capacidad representa el tamaño máximo del Array.
	// para agregar elementos a ese slice A , los primeros 5 (longitud) se 
	// pueden agregar como A[0] = "Alex" , luego de 5 si se debe usar append
	 nombres := make([]string, 5,10)
	 nombres[0] = "alex"
	 fmt.Println(nombres)

	 // para copiar de un slice a otro tambien se puede usar copy
	 // donde el primer elemento es donde se copiara el segundo elemento
	 // sintax : copy(slice1, slice2)
	 // copy devuelve solo el numero de elementos copiados
	 A := []int {1,2,3,4,5}
	 B := make([]int, 5) 
	 fmt.Println(B)
	 C := copy(B,A)
	 fmt.Println(B)
	 fmt.Println(C)

}