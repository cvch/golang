package main

import "fmt"

/*
Son interceptores que permiten ejecutar instrucciones comunes a varias funciones que reciben y devuleven
los mismos tipos de variables
*/
var result int

type operation func(int, int) int

/*
Realizar un ejemplo de interceptores en Go
*/
func main() {
	fmt.Println("inicio")

	result = operacionesMiddelware(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMiddelware(restar)(2, 7)
	fmt.Println(result)
	result = operacionesMiddelware(multiplicar)(4, 8)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMiddelware(funcion operation) operation {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")
		return funcion(a, b)
	}
}
