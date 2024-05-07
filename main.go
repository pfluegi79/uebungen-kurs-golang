package main

import (
	"fmt"
)

type ganzzahl int

var x ganzzahl
var y int

func main() {
	fmt.Printf("Der Wert von x ist: %v\n", x)
	fmt.Printf("Der Type von x ist: %T\n", x)
	x = 23
	fmt.Printf("Der neue Wert von x ist: %v\n", x)
	y = int(x)
	fmt.Printf("Der Wert von y ist: %v\n", y)
	fmt.Printf("Der Type von y ist: %T\n", y)
}
