package main

import "fmt"

func main() {

	/*
		[Entero con signo]
		- int : Depende de la arquitectura del sistema, normalmente 32 o 64 bits.
			ejemplo: -10, 0, 56
		- int8: Representa valores en el rango -128 a 127
		- int16: Representa valores en el rango -32768 a 32767
		- int32: -2147483648 a 2147483647
		- int64: -9223372036854775808 a 9223372036854775807

		[enteros sin signo]
		- uint8: entero cuyo rango es de 0 a 255
		- uint16: número entero cuyo rango va de 0 a 65535
		- uint32: entero cuyo rango es de 0 a 4294967295
		- uint64: entero cuyo rango es de 0 a 18446744073709551615

		[Conversión de tipos - Type Conversion]
		- Si convierte a un tipo que tiene un rango inferior al actual, se producirá una pérdida de datos.
			Hacemos typecast utilizando directamente el nombre de la variable como función para convertir tipos.

		[Desbordamiento - Overflow]
		- Si asignas un tipo y luego usas un número mayor que el rango del tipo para asignarlo, fallará.
			A continuación se muestra un programa que intenta precisamente eso.

	*/

	// Ejemplo básico
	var a int = 200 //int16
	var b int = -14 //int8
	sumResult := a + b
	fmt.Println("Result is: ", sumResult)

	// Type Conversion
	fmt.Println("\nType Conversion")
	var v1 int16 = 320
	var v2 int32 = int32(v1)
	fmt.Println("Value of v1 (int16): ", v1)
	fmt.Println("Value of v2 (int32): ", v2)

	// Overflow
	fmt.Println("\nOverflow")
	var maxInt8 int8 = 127
	var overflowInt8 int8 = maxInt8 + 1
	fmt.Println("Max. Int8 value: ", maxInt8)    // Valor esperado: 127
	fmt.Println("Int8 overflow: ", overflowInt8) // Salida esperada: Valor incorrecto o error interno.
}
