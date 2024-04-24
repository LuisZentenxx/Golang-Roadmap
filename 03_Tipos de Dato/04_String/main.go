package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		[Strings]
		- En Go, un string es una secuencia de caracteres Unicode codificados en UTF-8. Se utilizan ampliamente para representar texto y datos de texto.
		- Puedes crear un string de varias maneras, ya sea mediante una declaración explícita, inferencia de tipo o concatenación.
		- Los strings son inmutables en Go, lo que significa que no puedes modificar un carácter individual del string, pero puedes acceder a él para leerlo.
		- Puedes acceder a caracteres individuales de un string utilizando la indexación.

		[Funciones básicas integradas]
		- len(str): Devuelve la longitud del string
		- strings.ToLower(str): Devuelve una copia de un string con todas las letras en minusculas
		- strings.ToUpper(str): Devuelve una copia de un string con todas las letras en mayusculas
		- strings.Replace(str1, old, new, n) string: Devuelve una copia de str1 con n ocurrencias de old reemplazadas por new.

	*/

	// String vacio
	var emptyString string = ""
	fmt.Println(emptyString)

	// String con texto
	msg := "Hola, estoy programando en Golang"
	fmt.Println(msg)

	// String multilinea
	longMsg := `
	Esta es una cadena
	que ocupa multiples lineas
	en Golang.
	`
	fmt.Println(longMsg)

	// Concactenación de strings
	str1 := "Hola,"
	str2 := "mundo!"
	conStr := str1 + str2
	fmt.Println(conStr) // Salida esperada: "Hola, mundo!"

	// Longitud de un string
	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	longString := len(title)                               // Calcula la longitud del string
	fmt.Println("La longitud del titulo es: ", longString) // Salida esperada : "La longitud del titulo es: 55"

	// Convertir string en minusculas con toLower
	strUpper := "HOLA MUNDO"
	strLower := strings.ToLower(strUpper)
	fmt.Println(strLower) // Salida esperada: "hola mundo"

	// Convertir string en minusculas con toLower
	strMinus := "hola mundo"
	strMayus := strings.ToUpper(strMinus)
	fmt.Println(strMayus) // Salida esperada: "HOLA MUNDO"

	// Reemplazar n ocurrencias de old por new en s1
	s1 := "Hola, mundo! Esto es Golang!"
	old := "Hola"
	new := "Bienvenido"
	n := 1

	result := strings.Replace(s1, old, new, n)
	fmt.Println("Resultado: ", result)

}
