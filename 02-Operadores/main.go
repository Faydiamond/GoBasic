package main

import "fmt"

func main() {
	fmt.Println(" Operadores! ")

	age := 29

	fmt.Println(age < 30)
	//OR
	fmt.Println(age > 30 || age < 30)
	fmt.Println(age > 30 || age > 40)
	//AND
	fmt.Println(age < 30 && age > 0)
	//Not
	fmt.Println(!true)
	fmt.Println(!false)
}
