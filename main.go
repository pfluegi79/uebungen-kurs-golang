package main

import (
	"fmt"
)

var y = 42

var z string = "Flieht, ihr Narren!"

var a string = `Gandalf ruft: 
"Flieht, 
...
Ihr Narren"`

// Type bleibt bestehen und ist statisch (!) nicht wie in anderen Programmiersprachen dynamisch,
// also kann auch gerne mal wechseln

func main() {
	// Wertausgabe
	fmt.Println(y)
	// Typausgabe
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// z = 43 // Compiler Error
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)
}
