package main

import "fmt"

func main() {

	/*
		[Condicional Swtich-case]
		- La declaración switch-case permite evaluar una expresión y ejecutar diferentes bloques de código según el valor de esa expresión.
		- Es una alternativa más limpia y legible que múltiples declaraciones if-else anidadas en ciertos casos.

		[Uso]
		- Se evalúa la expresión y se compara con cada valor de caso.
		- Si se encuentra un caso que coincide con la expresión, se ejecuta el bloque de código asociado.
		- Si ninguno de los casos coincide, se ejecuta el bloque de código predeterminado (default) o simplemente se omite si no hay bloque predeterminado.

		[Notas]
		- No es necesario un "break" después de cada caso. Go automáticamente sale del switch después de ejecutar un caso.
		- Se pueden tener múltiples expresiones en un mismo caso si se separan con comas.
		- La expresión evaluada y los valores de caso deben ser del mismo tipo o convertibles entre sí.
	*/

	// Ejemplo básico
	option := 2

	// Utilizamos la declaración switch-case para manejar diferentes casos según el valor de "option".
	switch option {
	case 1:
		// Si "option" es igual a 1, se imprime un mensaje indicando que se seleccionó la opción 1: Abrir archivo.
		fmt.Println("Seleccionaste la opción 1: Abrir archivo")
	case 2:
		// Si "option" es igual a 2, se imprime un mensaje indicando que se seleccionó la opción 2: Guardar archivo.
		fmt.Println("Seleccionaste la opción 2: Guardar archivo")
	case 3:
		// Si "option" es igual a 3, se imprime un mensaje indicando que se seleccionó la opción 3: Cerrar archivo.
		fmt.Println("Seleccionaste la opción 3: Cerrar archivo")
	default:
		// Si "option" no coincide con ninguno de los casos anteriores, se ejecuta el caso predeterminado (default).
		// En este caso, se imprime un mensaje indicando que la opción no es válida.
		fmt.Println("Opción no válida")
	}

}
