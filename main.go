package main

import (
	"fmt"
)

var a int = 23
var b string = "Schlumpfine"
var c bool = true

func main() {
	s := fmt.Sprintf("%v %v %v", a, b, c)
	fmt.Println(s)
}
