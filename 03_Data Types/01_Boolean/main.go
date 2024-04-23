package main

import "fmt"

func main() {

	taskComplete := true
	testPassed := false

	award := taskComplete && testPassed

	fmt.Println("Student receives an award?: ", award) // Expected output : false

}
