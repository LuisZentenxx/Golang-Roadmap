package main

import "fmt"

func main() {

	// Variables númericas
	a := 20
	b := 10

	// --------------------------------------------------------
	// Utilizar operadores aritméticos
	sum := a + b
	subtraction := a - b
	multiplication := a * b
	division := a / b
	modulo := a % b

	// Imprimir los resultados
	fmt.Println("Sum:", sum)
	fmt.Println("Subtraction:", subtraction)
	fmt.Println("Multiplication:", multiplication)
	fmt.Println("Division:", division)
	fmt.Println("Modulus:", modulo)

	// --------------------------------------------------------

	// Utilizar operadores de comparación
	equal := (a == b)
	notEqual := (a != b)
	greaterThan := (a > b)
	lessThan := (a < b)
	greaterThanOrEqual := (a >= b)
	lessThanOrEqual := (a <= b)

	// Imrpimir los resultados
	fmt.Println("Equal:", equal)
	fmt.Println("Not equal:", notEqual)
	fmt.Println("Greater than:", greaterThan)
	fmt.Println("Less than:", lessThan)
	fmt.Println("Greater than or equal:", greaterThanOrEqual)
	fmt.Println("Less than or equal:", lessThanOrEqual)

	// --------------------------------------------------------

	// Usar operadores lógicos
	and := (a < 20) && (b > 2)
	or := (a < 20) || (b > 2)
	not := !(a == b)

	// Imprimir los resultados
	fmt.Println("Logical AND:", and)
	fmt.Println("Logical OR:", or)
	fmt.Println("Logical NOT:", not)

}
