package main

import (
	"ejer02/example"
	"encoding/json"
	"fmt"
	"strconv"
)

// variable que es exportada por que inicia en mayuscula
var Numero int
var texto string
var estatus bool

var numero8 int8
var numero16 int16

/*
Realice un ejemplo dónde se muestre el manejo del alcance (publicas y privadas) de las variables en Go
use una estructura
*/
func main() {
	var moneda int = 87
	// cast
	var flotante float32 = float32(moneda)

	texto = strconv.Itoa(int(moneda))

	var numero2, numero3 int
	numero2, numero3, texto, estatus := 5, 4, "texto", true
	fmt.Println(numero2, numero3, texto, estatus, moneda, flotante)
	mostrarEstatus()
	e := example.NewStruct(34, 56)
	jsonE, _ := json.Marshal(e)
	fmt.Println(string(jsonE))
}

func mostrarEstatus() {
	fmt.Println(estatus)
}
