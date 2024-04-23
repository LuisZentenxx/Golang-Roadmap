package main

import "fmt"

func main() {

	/*
		[Signed integers]
		- int : It depends on the architecture of the system, usually 32 or 64 bits.
			example: -10, 0, 56
		- int8: Represents values in the range -128 to 127
		- int16: Represents values in the range -32768 to 32767
		- int32: -2147483648 to 2147483647
		- int64: -9223372036854775808 to 9223372036854775807

		[Unsigned integers]
		- uint8: integer whose range is 0 to 255
		- uint16: integer whose range is 0 to 65535
		- uint32: integer whose range is 0 to 4294967295
		- uint64: integer whose range is 0 to 18446744073709551615

		[Type Conversion]
		- If you convert to a type that has range lower than your current range, data loss will occur.
			We do typecast by directly using the name of the variable as a function to convert types.

		[Overflow]
		- If you assign a type and then use a number larger than the types range to assign it, it will fail.
			Below is a program trying just that.

	*/

	// Basic example
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
	fmt.Println("Max. Int8 value: ", maxInt8)    // Expected value: 127
	fmt.Println("Int8 overflow: ", overflowInt8) // Expected output: Incorrect value or error
}
