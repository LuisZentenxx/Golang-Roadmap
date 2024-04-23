package main

import "fmt"

func main() {
	/*
		- float32: Representa números con hasta aproximadamente 7 dígitos decimales de precisión.
		- float64: Representa números con una precisión de hasta 15 dígitos decimales aproximadamente.
	*/

	var x float32 = 12.1415953         // Decimal de 7 digitos
	var y float64 = 42.560914592840425 // Decimal de 15 digitos
	fmt.Println("El valor de x (float32) es: ", x)
	fmt.Println("El valor de y (float32) es: ", y)

	var z float32 = float32(y)
	fmt.Println(z) // Sslida esperada: Valor truncado o redondeado con imprecisión o valor incorrecto.

}
