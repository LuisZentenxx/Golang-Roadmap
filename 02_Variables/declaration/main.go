package main

// Import 'fmt' package providing functions for input and formatted output
import "fmt"

func main() {

	/*
		Variable Declaration
		-----------------------
			- var (Keyword)
			- name (variable name)
				Correct: variable name in CamelCase, first letter lowercase and following words uppercase: (var nameLanguage string)
				Incorrect: variable name with all words capitalized: (var NameLanguage string)
				Incorrect: variable name with underscore to separate words: (var Name_Language string)
			- type (data type)
			- value (variable value) [Optional]
	*/

	// variable not initialized
	var age int

	// Initialize the variable at the same time as you declare it
	var nameLanguage string = "Golang"

	// Print the variable name
	fmt.Println("This program is made in ", nameLanguage)

}
