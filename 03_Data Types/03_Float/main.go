package main

import "fmt"

func main() {
	/*
		- float32: Represents numbers with up to approximately 7 decimal digits of accuracy.
		- float64: Represents numbers with up to approximately 15 decimal digits of accuracy.
	*/

	var x float32 = 12.1415953         // 7 decimal digits
	var y float64 = 42.560914592840425 // 15 decimal digits
	fmt.Println("Value of x (float32) is: ", x)
	fmt.Println("Value of y (float32) is: ", y)

	var z float32 = float32(y)
	fmt.Println(z) // Expected output: truncated or rounded and imprecise 'y' value

}
