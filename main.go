package main

import (
	"fmt"
)

// Deklaration der Variable "y"
// Wertzuweisung ist 23
// Implizierte Typzuweisung und damit Initialisierung
// Deklaration UND Wertzuweisung = Initialisierung
var y = 23

// Deklaration der Variable "z"
// Typzuweisung: Identifier "z" ist Typ int
// Implizierte Wertezuweisung und Initialisierung durch automatische Zuweisung
// eines "Null"-Wertes (in manchen Fällen NIL)
// d.h. z. B.: false für booleans, 0 für integers, 0.0 für floats, "" für strings,
// und NIL für pointers, functions, interfaces, slices, channels und maps.
var z int

func main() {
	// Short Declaration Operator
	// Deklaration UND Wertzuweisung (eines implizierten Typs)
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("y ist global nutzbar sogar in Unterfunktionen:", y)
}
