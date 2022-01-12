package main

import "fmt"

var status bool

/*
Realice un ejemplo de las estructuras de control if, else if, switch
y el manejo de asiganciones antes de ejecutar la clausula
*/
func main() {
	status = true
	if status = false; status == true {
		fmt.Println("estado: true", status)
	} else {
		fmt.Println("estado: false", status)
	}

	var number int
	if number == 2 {
		fmt.Println("Numero Es igual a 2: ", number)
	} else if number = 4; number == 3 {
		fmt.Println("Numero Es igual a 3: ", number)
	} else {
		fmt.Println("Numero Es igual a? valor: ", number)
	}

	switch numberCase := 5; numberCase {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	default:
		fmt.Println("Acción por defecto: ", numberCase)
	}
	switch numberCase := 5; {
	case numberCase > 1:
		fmt.Println("1")
		fallthrough
	case numberCase > 2:
		fmt.Println("2")

	default:
		fmt.Println("Acción por defecto: ", numberCase)
	}

}
