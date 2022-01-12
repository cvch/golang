package main

import (
	"ejer009/user"
	"fmt"
	"time"
)

/*
Definir una estructura que haga uso de la composición usando modulos
*/
type pepe struct {
	user.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Pepe Ramirez", time.Now(), true)
	fmt.Println(u.Usuario)
}
