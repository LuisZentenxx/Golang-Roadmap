package main

import "fmt"

func main() {
	/*

	 */

	// Empty string
	var emptyString string = ""
	fmt.Println(emptyString)

	// String with text
	msg := "Hi, I'm programming in Golang"
	fmt.Println(msg)

	// Multiline string
	longMsg := `
	This is a string
	that occupies multiple lines
	in Go.
	`
	fmt.Println(longMsg)

	// String concatenation
	str1 := "Hello,"
	str2 := "world!"
	conStr := str1 + str2
	fmt.Println(conStr) // Expected outpud: "Hello, world!"

}
