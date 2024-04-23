package main

import "fmt"

func main() {

	taskComplete := true
	testPassed := false

	award := taskComplete && testPassed

	fmt.Println("¿El estudiante recibe un premio?: ", award) // Salida esperada: false

}
