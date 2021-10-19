package main

import (
	"fmt"
)

func main() {
	//paises := make(map[string]string, 5) puedo reservar espacio

	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenmos Aires"

	fmt.Println(paises["Mexico"])
	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30}
	campeonato["Millos"] = 25
	campeonato["Naciomal"] = 0
	delete(campeonato, "Naciomal")
	fmt.Println(campeonato)
	campeonato["Millos"] = 80

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d\n", equipo, puntaje)
	}

	//Existe en un mapa
	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
