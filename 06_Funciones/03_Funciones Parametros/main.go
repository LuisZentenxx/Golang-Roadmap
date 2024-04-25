package main

import "fmt"

/*
[Funciones con Parámetros y Retorno]
	- Las funciones con parámetros y retorno son bloques de código que toman ciertos valores como entrada (parámetros), realizan una tarea específica y luego devuelven un resultado como salida.
	- Estas funciones son útiles cuando necesitas realizar un cálculo o procesamiento basado en datos proporcionados como entrada y obtener un resultado específico como salida.

- Sintaxis:
     La sintaxis básica para definir una función en Go que tome parámetros y devuelva un valor es la siguiente:
	 	func nombreDeLaFuncion(parametro1 tipo1, parametro2 tipo2, ...) tipoDeRetorno {
			// Cuerpo de la función
		}

- Llamada a la función:
     Una vez que una función está definida, puede ser llamada desde cualquier parte del programa utilizando su nombre seguido de paréntesis y proporcionando valores para los parámetros esperados. Por ejemplo:
     resultado := nombreDeLaFuncion(valor1, valor2, ...)

- Retorno:
     Para devolver un valor desde una función, se utiliza la palabra clave `return` seguida del valor que deseas devolver. Por ejemplo:
     return valor


[Funciones con Parámetros y Sin Retorno]
	- Las funciones con parámetros y sin retorno son bloques de código que toman ciertos valores como entrada (parámetros) pero no devuelven ningún valor como resultado.
	- Estas funciones son útiles cuando necesitas realizar alguna tarea o acción basada en los parámetros proporcionados, pero no necesitas que la función devuelva ningún resultado.

- Sintaxis:
  La sintaxis básica para definir una función en Go que tome parámetros pero no devuelva ningún valor es la misma que para las funciones con retorno, pero sin especificar el tipo de retorno:
  	func nombreDeLaFuncion(parametro1 tipo1, parametro2 tipo2, ...) {
		// Cuerpo de la función
	}

- Llamada a la función:
  Una vez que una función está definida, puede ser llamada desde cualquier parte del programa utilizando su nombre seguido de paréntesis y proporcionando valores para los parámetros esperados.
  Por ejemplo: nombreDeLaFuncion(valor1, valor2, ...)

*/

// Función que calcula el área de un rectángulo
func calculateRectangleArea(base, height float64) float64 {
	return base * height
}

// Función que calcula el promedio de una lista de números
func calculateAverage(numbers []int) float64 {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return float64(sum) / float64(len(numbers))
}

// Función con parametro y sin retorno que imprime un saludo personalizado
func printGreeting(greeting string) {
	fmt.Println("Hola ", greeting)
}

func main() {

	// Llamamos a la función calculateRectangleArea con los valores de base y altura
	area := calculateRectangleArea(4.3, 2.2)
	fmt.Println("El área del rectangulo es: ", area)

	// Lista de números
	numbersList := []int{43, 20, 92, 23, 12}

	// Llamamos a la función calculateAverage con la lista de números como argumento
	average := calculateAverage(numbersList)
	fmt.Println("El promedio de la lista de números es: ", average)

	// Llamamos a la función printGreeting con el parámetro "Mundo"
	printGreeting("Mundo")
}
