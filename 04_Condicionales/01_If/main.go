package main

import "fmt"

func main() {

	/*
		[Condicional If-else]
			- If: Permite ejecutar un bloque de código si se cumple una condición
			- Else: Opcionalmente permite ejecutar un bloque de código alternativo si la condición no se cumple.

		[Condicional If Anidado]
		    - El condicional "if" anidado es una estructura en la que se coloca un "if" dentro de otro "if".
		    - Si la primera condición evaluada por el primer "if" es falsa, se verifica la condición del "if" anidado.
		    - Esto permite manejar situaciones donde se necesita evaluar múltiples condiciones de manera secuencial.
	*/

	// Ejemplo If basico
	edad := 18

	// Si la variable "edad" es mayor o igual a 18, se imprime un mensaje.
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")

		// En caso contrario, se imprime otro mensaje.
	} else {
		fmt.Println("Eres menor de edad")
	}

	// Ejemplo de una estructura if/else anidada en Go

	num := 30

	// En este ejemplo, estamos evaluando la variable "num" para determinar si es mayor que 0, menor que 0 o igual a 0.
	if num > 0 {
		// Si "num" es mayor que 0, se imprime un mensaje indicando que es mayor a 0.
		fmt.Println("El número es mayor a 0")
	} else if num < 0 {
		// Si "num" es menor que 0, se imprime un mensaje indicando que es menor a 0.
		fmt.Println("El número es menor a 0")
	} else {
		// Si "num" no es mayor que 0 ni menor que 0, entonces deducimos que es igual a 0 y se imprime el mensaje correspondiente.
		fmt.Println("El número es igual a 0")
	}

}
