package main

import (
	"fmt"
	"math" // Se importa la biblioteca `math` que permite trabajar con operaciones matematicas más complejas
)

func main() {

	/*
		[Números complejos]
		- complex: Esta función se utiliza para crear un número complejo a partir de sus partes real e imaginaria.
			Por ejemplo: complex(2, 3) crea el número complejo 2+3𝑖
		- real(complex): Devuelve la parte real de un número complejo.
		- imag(complex): Devuelve la parte imaginaria de un número complejo.
		- conj(complex): Devuelve el conjugado de un número complejo.
			El conjugado de 𝑎+𝑏𝑖 es 𝑎-b𝑖
		- abs(complex): Calcula el módulo (o valor absoluto) de un número complejo. Es decir, devuelve la distancia del número complejo al origen en el plano complejo.
		- phase(complex): Calcula el ángulo (o fase) de un número complejo. Devuelve el ángulo en radianes entre el eje x positivo y la línea que une el origen
			con el punto representado por el número complejo en el plano complejo.
		- polar(complex) (float64, float64): Convierte un número complejo en coordenadas polares y devuelve el módulo y la fase.
		- rect(float64, float64) complex128: Convierte coordenadas polares (módulo y fase) en un número complejo en coordenadas cartesianas.
	*/

	// Ejercicio básico

	// La variable "num" almacena un número complejo a partir de sus partes real e imaginaria (3,4) respectivamente.
	num := complex(3, 4)

	// La variable "modulus" almacena el módulo del número complejo. La función `math.Sqrt()` calcula la raíz cuadrada.
	modulus := math.Sqrt(real(num)*real(num) + imag(num)*imag(num))

	fmt.Println("El módulo es:", modulus)
}
