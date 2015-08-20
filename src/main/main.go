package main

import (
	"fmt"
	"utility"
	"strings"	
)

func main() {
	
	// First Example
	x := 10
	x++
	fmt.Println("Hello", x)
	
	// Struct
	p := utility.Persona{Nome: "Mario", Cognome: "Rossi"}
	p.Stampa()	
	p1 := utility.CreatePersona("Giuseppe", "Bianchi")
	p1.Stampa()
	
	// Interface
	q := utility.Quadrato{X: 10, Y: 15}
	f := utility.FormaGeometrica(q)
	a := CalcolaArea(f)
	//a := CalcolaArea(&q)
	fmt.Println("Area:", a)
	
	// Array (static)
	vs := [2]string{}
	vs[0] = "Mario"
	vs[1] = "Bruno"
	for _, v := range vs {
		fmt.Println(v)
	}
	
	// Array (dynamic)
	vd := []string{}
	vd = append(vd, "Giulio")
	vd = append(vd, "Maria")
	for _, v := range vd {
		fmt.Println(v)
	}
	
	// Map
	m := map[string]string{}
	m["k1"] = "Pippo"
	m["k2"] = "Pluto"
	for k, v := range m {
		fmt.Println(k, "=", v)
	}
	
	// Closure
	result := compareint(10, 20, func(a int, b int) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			} else {
				return 0
			}	 
	});
	fmt.Println("Result: ", result)
	
	// Transformer Map
	tm := map[string]func(string) string{}
	tm["t1"] = func(k string) string {
		return strings.ToUpper(k)
	}
	fmt.Println("Transformer: ", tm["t1"]("mario"))	
}

func CalcolaArea(forma utility.FormaGeometrica) int {
	return forma.Area()
}

func compareint(a int, b int, mycomparator func(int, int) int) int {
	return mycomparator(a, b)
}

func dummy(k string) string {
	return strings.ToLower(k)
}