package main

import (
	"fmt"
	"time"
)

/*
Realizar un ejercicio que llame una función en paralelo y reciba el valor retornado por esta
*/
func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Legué hasta aquí")

	msg := <-canal1
	fmt.Println(msg)
}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 10000000000; i++ {

	}

	final := time.Now()
	canal1 <- final.Sub(inicio)
}
