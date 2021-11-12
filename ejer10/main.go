package main

import "fmt"

type SerVivo interface {
	EstaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	EstaVivo() bool
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	EstaVivo() bool
}

type vegetal interface {
	ClasificacionVegetal() string
	EstaVivo()
}

/* Genero Humano */
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHomre    bool
	vivo       bool
}

type mujer struct {
	hombre
	// edad       int
	// altura     float32
	// peso       float32
	// respirando bool
	// pensando   bool
	// comiendo   bool
	// esHomre    bool
	// vivo       bool
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHomre {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) EstaVivo() bool { return h.vivo }

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, ya ya estoy respirando \n", hu.sexo())
}

/*Genero Animal*/

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }
func (p *perro) EstaVivo() bool    { return p.vivo }

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() {
		return 1
	}
	return 0
}

func estoyVivo(v SerVivo) bool {
	return v.EstaVivo()
}

func main() {
	Pedro := new(hombre)
	Pedro.esHomre = true
	HumanosRespirando(Pedro)

	Maria := new(mujer)
	HumanosRespirando(Maria)

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros = +AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoros %d \n", totalCarnivoros)
	fmt.Printf("Estoy vivo = %t \n", estoyVivo(Dogo))
	fmt.Printf("Estoy vivo = %t \n", estoyVivo(Maria))

}
