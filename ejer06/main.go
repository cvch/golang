package main

import (
	"fmt"
)

func main() {
	fmt.Println(uno(7))
	fmt.Println(dos(1))
	fmt.Println(dos(7))
	var aux1 int
	var aux2 bool
	_, aux2 = dos(7)
	fmt.Println("valor de aux1 y de aux2 ", aux1, aux2)
	fmt.Println(dos(7))
	println("---calculo---")
	fmt.Println(calculo(7, 1, 2, 3, 4, 5))
	// arr := []int{7, 1, 2, 3, 4, 5} //lo mismo que lo anterior
	// fmt.Println(calculo(arr...))
}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func calculo(numero ...int) int {
	total := 0
	/*
		la _ para go hace que el leguaje no reserve ni separe ni use memoria
		para el valor devuelto por la instruccion range
		for _, num := range numero
	*/
	for indice, num := range numero {
		total = total + num
		fmt.Printf("indice %d \n", indice)
		fmt.Printf("num %d \n", num)
		fmt.Printf("total %d \n\n", total)
	}
	return total
}
