package main

// Importar el paquete 'fmt' que proporciona funciones para la entrada y la salida formateada.
import "fmt"

func main() {

	// Variable no inicializada al momento de declararla. (Se inicializa en la siguiente linea otorgando un valor númerico (25))
	var age int
	age = 25
	fmt.Println("Mi edad es ", age) // Salida esperada: "Mi edad es 25"

	// Variable inicializada al mismo tiempo que se declara
	var nameLanguage string = "Golang"

	// Imprime el valor de la variable
	fmt.Println("Este programa está hecho en Golang ", nameLanguage) // Salida esperada: "Este programa está hecho en Golang"

	// Variable declarada e inicializada al mismo tiempo usando el operador ":="
	price := 3000

	fmt.Println("El precio es ", price) // Salida esperada: "El precio es 3000".

}
