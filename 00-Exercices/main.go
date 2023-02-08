package main

import "fmt"

func main() {
	/* Variables */
	var name1 string
	name1 = "Fabian"

	fmt.Printf("Su nombre es: %s \n", name1)
	var age int
	age = 29
	fmt.Printf("Su edad es : %d \n", age)

	/* Variables inferidas  */
	var mobiles [5]string
	mobiles[0] = "Samsung"
	mobiles[1] = "Iphone"
	mobiles[2] = "Motorola"
	mobiles[3] = "Xiaomi"
	mobiles[4] = "Honor"

	for _, mark := range mobiles {
		fmt.Println(mark)
	}

	bags := []string{"Wilson", "Velez", "Bosi", "Totto", "Arturo calle"}

	for i := 0; i < len(bags); i++ {
		fmt.Printf("Esto es i %d valor %s", i, bags[i])
	}
}
