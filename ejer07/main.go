package main

import (
	"fmt"
)

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5, 7))

	// Restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 4 = %d \n ", Calculo(6, 4))

	//Dividimos
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Dividimos 8 / 2 = %d \n ", Calculo(8, 2))

	Operaciones()

	/* CLOSURES */
	tablaDel := 3
	MiTabla := Tabla(tablaDel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
	inputSimultion := []string{"esto", "es", "un", "texto", "de", "un", "archivo"}
	myScanner := scannerSimulation(inputSimultion)
	for i := 0; i < len(inputSimultion); i++ {
		fmt.Println(myScanner())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}

	fmt.Println(resultado())
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}

func scannerSimulation(input []string) func() string {
	index := 0
	return func() (text string) {
		text = input[index]
		index++
		return
	}
}
