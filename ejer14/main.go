package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	// con decirle go ya me inicia asincronicamente la funcion

	wg.Add(1)
	go miNombreLentooo("Camilo Velasco Chaves", &wg)
	fmt.Println("Estoy Aqui")
	wg.Wait()
}

func miNombreLentooo(nombre string, wg *sync.WaitGroup) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
	wg.Done()
}
