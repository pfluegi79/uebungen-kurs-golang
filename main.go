package main

import (
	"fmt"
)

var y int

func main() {
	fmt.Println("Hello World!")
	fmt.Printf("Hello meine Variable heisst y. %v\n", y)
	y = 23
	fmt.Printf("%v\n", y)
	fmt.Printf("%#x\t %X\n", y, y)
	fmt.Printf("%b\n", y)
	fmt.Print(y)
	fmt.Print(y)
	fmt.Println(y)
	ausgabe := fmt.Sprintf("\n%#x\t %X\n", y, y)
	fmt.Println(ausgabe)

}
