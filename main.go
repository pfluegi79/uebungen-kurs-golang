package main

import "fmt"

func main() {
	fmt.Println("Hallo User!!")
	foo()
	fmt.Println("Foo ist beendet. Hier geht's weiter im Hauptteil")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("Hier sind wir in der Funktion Foo()")
}

func bar() {
	fmt.Println("Hier sind wir in der Funktion Bar()")
}
