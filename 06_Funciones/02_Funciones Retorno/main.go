package main

import "fmt"

/*
[Funciones con Retorno y Sin Parámetros]
	Las funciones con retorno y sin parámetros son bloques de código que realizan una tarea específica cuando son llamadas, pero no requieren ningún valor de entrada para hacer su trabajo.
		Sin embargo, pueden devolver un valor como resultado de su ejecución.

[Para qué sirven]
   Estas funciones son útiles cuando necesitas realizar una tarea que no depende de ningún dato específico, pero quieres obtener un resultado al final.

- Sintaxis:
	La sintaxis básica para definir una función en Go que tenga retorno pero sin parámetros es la siguiente:
		func nombreDeLaFuncion() tipoDeRetorno {
         	// Cuerpo de la función
		}
- Uso:
     Una vez que una función está definida, puede ser llamada desde cualquier parte del programa utilizando su nombre seguido de paréntesis (). Por ejemplo:
     nombreDeLaFuncion()

- Retorno:
     Para devolver un valor desde una función, se utiliza la palabra clave `return` seguida del valor que deseas devolver. Por ejemplo:
     return valor.
*/

// Función que devuelve un mensaje
func getMessage() string {
	return "Esto es un ejemplo de funcion con retorno en Golang"
}

// Función que devuelve el nombre de un lenguaje de programación según el indice
func getProgrammingLangugage() string {
	progLanguages := []string{"Golang", "Python", "Java", "Kotlin", "Haskell"}
	index := 2
	return progLanguages[index]
}

func main() {

	msg := getMessage()
	fmt.Println(msg)

	// lLamamos a la funcion getProgrammingLangugage y mostramos su resultado
	output := getProgrammingLangugage()
	fmt.Println(output) // Salida Esperada: "Java"
}
