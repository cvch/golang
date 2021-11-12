package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	archivo := "prueba.txt"
	f, err := os.Open(archivo)

	// se ejecuta siempre antes de salir de la función
	defer f.Close()

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		//os.Exit(1)
	}

	ejemploPanic()
}

func ejemploPanic() {
	defer func() {
		// recover funcion que va y mira si ocurrio un panic
		r := recover()
		if r != nil {
			log.Fatalf("Ocurrió un error que generó un panic \n %v", r)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor de 1")
	}
}
