package main

import (
	"bufio"
	"fmt"
	"os"
)

var number1 int
var number2 int
var text string
var result int

func main() {
	fmt.Println("Ingrese número 1: ")
	fmt.Scanln(&number1)
	fmt.Println("Ingrese número 2: ")
	fmt.Scanln(&number2)
	fmt.Println("La suma de los números es: ", number1+number2)
	fmt.Println("Ingrese descripción: ")

	// se puede hacer todo con este scanner y es mas rapido pero
	// toca hacer todo manual, como el parseo a enteros, etc.
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}
	result = number1 + number2
	fmt.Println(text, result)
}
