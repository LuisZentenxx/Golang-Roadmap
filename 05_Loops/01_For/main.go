package main

import "fmt"

func main() {
	/*
		[For Loop]
			- El bucle 'for' es una estructura de control utilizada en Go para repetir un bloque de código múltiples veces.

			- Sintaxis:
				for inicialización; condición; post-acción {
					// codigo a ejecutar
				}

		   - La 'inicialización' se ejecuta antes de la primera iteración del bucle y generalmente se utiliza para inicializar variables de control.

		   - La 'condición' se evalúa antes de cada iteración. Si es verdadera, el bloque de código dentro del bucle se ejecuta. Si es falsa, el bucle se detiene.

		   - La 'post-acción' se ejecuta después de cada iteración y generalmente se utiliza para actualizar variables de control.

		   - Ejemplo:
		     for i := 0; i < 5; i++ {
		         // código a ejecutar
		     }
		     Este bucle imprimirá los números del 0 al 4.
	*/

	// Uso de for para imprimir los números del 1 al 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Uso de for para imprimir los números impares del 1 al 10
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// Inicializamos una variable suma para almacenar la suma de los números
	suma := 0

	// Utilizamos un bucle 'for' para sumar los números del 1 al 10
	for i := 1; i <= 10; i++ {
		suma += i
	}

	// Imprimimos la suma
	fmt.Println("La suma de los números del 1 al 10 es:", suma)

}
