package main

import "fmt"

func main() {
	/*
		[Range Loop]
			- El bucle 'range' se utiliza en Go para iterar sobre elementos de una colección como arrays, slices, mapas, o canales.

		   - Sintaxis:
		     for indice, valor := range colección {
		         // código a ejecutar
		     }

		   - 'indice' representa el índice (o clave en el caso de mapas) del elemento actual en la colección.

		   - 'valor' representa el valor del elemento actual en la colección.

		   - Ejemplo:
		     nombres := []string{"Juan", "María", "Carlos"}
		     for indice, nombre := range nombres {
		         fmt.Println(indice, nombre)
		     }
		     Este bucle imprimirá los índices y los nombres en el slice 'nombres'.
	*/

	nombres := []string{"Juan", "María", "Carlos"}
	for indice, nombre := range nombres {
		fmt.Println(indice, nombre)
	}

	texto := "Golang"

	// Utilizamos el bucle 'range' para iterar sobre la cadena de texto 'texto'
	for _, caracter := range texto {
		fmt.Println(string(caracter))
	}

}
