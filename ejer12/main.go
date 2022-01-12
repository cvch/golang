package main

import "fmt"

/*
Realice un ejemplo de recursividad
*/
func main() {
	exponencia(2)
}

func exponencia(numero int) {
	if numero > 100000000 {
		return
	}
	fmt.Println(numero)
	exponencia(numero * 2)
}
