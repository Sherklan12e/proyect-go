package main

import (
	"fmt"
)

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) Saludar() {
	fmt.Printf("hola me llamo %s y tengo %d \n", p.Nombre, p.Edad)
}
func (p *Persona) Cumpliranito() {
	p.Edad++
	fmt.Printf("ya tienes %d \n", p.Edad)
}
func main() {
	p1 := Persona{"juan", 20}
	p1.Saludar()
	p1.Cumpliranito()

	p1.Cumpliranito()
	p1.Cumpliranito()
}
