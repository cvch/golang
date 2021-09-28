package main

import (
	"fmt"
	"strconv"
)

// variable que es exportada por que inicia en mayuscula
var Numero int
var texto string
var estatus bool

var numero8 int8
var numero16 int16

func main() {
	var moneda int = 87
	// cast
	var flotante float32 = float32(moneda)

	texto = strconv.Itoa(int(moneda))

	var numero2, numero3 int
	numero2, numero3, texto, estatus := 5, 4, "texto", true
	fmt.Println(numero2, numero3, texto, estatus, moneda, flotante)
	mostrarEstatus()
}

func mostrarEstatus() {
	fmt.Println(estatus)
}
