package main

import "fmt"

func main() {
	n, _ := fmt.Println("Was ist der Sinn des Lebens?", 42, true)
	fmt.Println(n)
	n2, err := fmt.Println("Was ist der Sinn des Lebens?", 42, true)
	fmt.Println(n2)
	fmt.Println(err)
}
