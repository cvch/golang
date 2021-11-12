package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// con decirle go ya me inicia asincronicamente la funcion
	miNombreLentooo("Camilo Velasco Chaves")
	fmt.Println("Estoy Aqui")
	var x string
	fmt.Scanln(&x)
}

func miNombreLentooo(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
