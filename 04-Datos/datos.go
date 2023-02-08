package main

import "fmt"

func main() {
	//definir un unico caracter comillas simples
	var A byte = 'A'
	var a byte = 'a'
	var R byte = 82
	var s byte = 115

	fmt.Println()
	fmt.Println("Value  to ASCII Code:")
	fmt.Println(A)
	fmt.Println(a)
	fmt.Println(R)
	fmt.Println(s)

	fmt.Println()
	fmt.Println("Value  to ASCII Code:")
	fmt.Println(string(A))
	fmt.Println(string(a))
	fmt.Println(string(R))
	fmt.Println(string(s))
}
