package main

import (
	"fmt"
)

var a int
var b ganzzahl

type ganzzahl int

func main() {
	a = 23
	fmt.Printf("a hat den Wert:  %v\n", a)
	fmt.Printf("a hat den Typ:  %b\n", a)
	fmt.Printf("a hat den Typ:  %T\n", a)
	b = 42
	fmt.Printf("a hat den Wert:  %v\n", b)
	fmt.Printf("a hat den Typ:  %b\n", b)
	fmt.Printf("a hat den Typ:  %T\n", b)
	//funktioniert nicht
	// a = b
}
