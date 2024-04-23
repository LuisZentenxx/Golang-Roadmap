package main

import "fmt"

func main() {

	/*
		Operadores
		------------
		[Operadores aritméticos]
		- (+) equivale a suma
		- (-) equivale a la resta
		- (*) equivale a la multiplicación
		- (/) equivale a la división
		- (%) equivale al resto de una operación de división entre dos números. (módulo)

		[Operadores de comparación]
		- (==) equivale a igual
		- (!=) equivale a diferente de
		- (>) equivale a mayor que
		- (<) equivale a menor que
		- (>=) equivale a mayor o igual que
		- (<=) equivale a menor o igual que

		[Operadores lógicos]
		- (&&) es equivalente a Y (AND) -> (a < 20) && (b > 2)
		- (||) es equivalente a O (OR) -> (a < 20) || (b > 2)
		- (!) es equivalente a NO (NOT) -> !(a == b)
	*/

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
