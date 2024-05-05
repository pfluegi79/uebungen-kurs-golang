package main

import (
	"fmt"
)

type ganzzahl int

var x ganzzahl

func main() {
	fmt.Printf("Der Wert von x ist: %v\n", x)
	fmt.Printf("Der Type von x ist: %T\n", x)
	x = 23
	fmt.Printf("Der neue Wert von x ist: %v\n", x)
}
