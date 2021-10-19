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

	var iterator = 0

	for iterator < 10 {
		fmt.Printf("\n Valor de iterator: %d", iterator)
		if iterator == 5 {
			fmt.Printf("  Sumamos 1 \n")
			iterator = iterator + 1
			continue
		}
		fmt.Printf("   Entra aquí \n")
		iterator++
	}

	var ite = 0
BLOQUE1:
	fmt.Println("Inicio del Bloque1")
	for ite < 10 {
		fmt.Println("valor de ite: ", ite)
		if ite == 5 {
			ite = ite + 1
			fmt.Println("Ite -1 y voy a BLOQUE1")
			goto BLOQUE1
		}
		fmt.Println("Incremento en 1 a ite")
		ite++
	}

}
