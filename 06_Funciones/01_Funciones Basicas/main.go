package main

import "fmt"

/*
[Funciones Básicas]
-Las funciones son bloques de código que realizan una tarea específica cuando son llamadas.
	En Go, las funciones pueden ser definidas y luego invocadas en cualquier parte del programa.

- Sinraxis:
	La sintaxis básica para definir una función en Go es la siguiente:
		func nombreDeLaFuncion() {
       		// Cuerpo de la función
		}
- Uso:
	Una vez que una función está definida, puede ser llamada desde cualquier parte del programa utilizando su nombre seguido de paréntesis ().
	Por ejemplo: nombreDeLaFuncion()
*/

// Función básica sin parámetros ni retorno
func greetings() {
	fmt.Println("Hola, mundo!")
}

// Función básica que hace una operación matematica y entrega un resultado
func sumDigits() {
	z := 3
	x := 42.3
	y := -43

	result := z + int(x) + y
	fmt.Println("Resultado: ", result)

}

func main() {

	// Se llama a la función greetings
	greetings()

	// Se llama a la función sumDigits
	sumDigits()

}
