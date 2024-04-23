package main

import "fmt"

func main() {

	/*
		Operators
		------------
		[Arithmetic operators]
		- (+) is equivalent to sum
		- (-) is equivalent to subtraction
		- (*) is equivalent to multiplication
		- (/) is equivalent to division
		- (%) is equivalent to remainder of a division operation between two numbers. (modulus)

		[Comparison operators]
		- (==) is equivalent to equals
		- (!=) is equivalent to not equals
		- (>) is equivalent to greater than
		- (<) is equivalent to less than
		- (>=) is equivalent to greater than or equal to
		- (<=) is equivalent to less than or equal to

		[Logical operators]
		- (&&) is equivalent to and -> (a < 20) && (b > 2)
		- (||) is equivalent to or -> (a < 20) || (b > 2)
		- (!) is equivalent to not -> !(a == b)
	*/

	// Declare two numeric variables
	a := 20
	b := 10

	// --------------------------------------------------------
	// Use arithmetic operators
	sum := a + b
	subtraction := a - b
	multiplication := a * b
	division := a / b
	modulo := a % b

	// Print the results
	fmt.Println("Sum:", sum)
	fmt.Println("Subtraction:", subtraction)
	fmt.Println("Multiplication:", multiplication)
	fmt.Println("Division:", division)
	fmt.Println("Modulus:", modulo)

	// --------------------------------------------------------

	// Use comparison operators
	equal := (a == b)
	notEqual := (a != b)
	greaterThan := (a > b)
	lessThan := (a < b)
	greaterThanOrEqual := (a >= b)
	lessThanOrEqual := (a <= b)

	// Print the results
	fmt.Println("Equal:", equal)
	fmt.Println("Not equal:", notEqual)
	fmt.Println("Greater than:", greaterThan)
	fmt.Println("Less than:", lessThan)
	fmt.Println("Greater than or equal:", greaterThanOrEqual)
	fmt.Println("Less than or equal:", lessThanOrEqual)

	// --------------------------------------------------------

	// Use logical operators
	and := (a < 20) && (b > 2)
	or := (a < 20) || (b > 2)
	not := !(a == b)

	// Print the results
	fmt.Println("Logical AND:", and)
	fmt.Println("Logical OR:", or)
	fmt.Println("Logical NOT:", not)

}
