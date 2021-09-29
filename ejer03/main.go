package main

import "fmt"

var status bool

func main() {
	status = true
	if status = false; status == true {
		fmt.Println("estado: ", status)
	} else {
		fmt.Println("estado: ", status)
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
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("Acción por defecto: ", numberCase)
	}

}
