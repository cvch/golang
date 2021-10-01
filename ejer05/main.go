package main

import "fmt"

func main() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for ii := 1; ii <= 10; ii++ {
		fmt.Println(ii)
	}

	number := 0
	for {
		fmt.Println("Ciclo infinito mientras se cumpla la condición")
		fmt.Println("ingrese el número 1 para detener el ciclo")
		fmt.Scanln(&number)
		if number == 1 {
			break
		}
	}
}
