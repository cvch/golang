package main

/*
Definir una estructura que haga uso de la composición
*/
import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct {
	us.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Pepe Ramirez", time.Now(), true)
	fmt.Println(u.Usuario)
}
