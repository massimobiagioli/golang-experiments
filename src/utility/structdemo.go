package utility

import (
	"fmt"
)

type Persona struct {
	Nome string
	Cognome string
}

func CreatePersona(n string, c string) (*Persona) {
	p := Persona {Nome: n, Cognome: c}
	return &p
}

func (p Persona) Stampa() {	
	fmt.Println(p.Nome + " " + p.Cognome)
}
