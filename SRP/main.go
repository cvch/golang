/*
 package main
 2
 3 type Documento struct {
 4     Nombre  string
 5 }
 6
 7 func (d *Documento) GuardarEnArchivo() {
 8    // ...
 9 }
10
11 func (d *Documento) GuardarEnBaseDatos() {
12    // ...
13 }
*/

package main

import "fmt"

type PersistirDocumento interface {
	Guardar(d Documento)
}

type Documento struct {
	Nombre string
}

type DocumentoBD struct {
	Documento
	StringConnection string
}

func (documento *Documento) Guardar() {
	fmt.Println("Guardo en archivo local")
}

func (documentoBD *DocumentoBD) Guardar() {
	fmt.Println("Guardo en archivo BD")
}

func main() {
	documento := new(Documento)
	documento.Guardar()

	documentoBD := new(DocumentoBD)
	documentoBD.Guardar()
}
