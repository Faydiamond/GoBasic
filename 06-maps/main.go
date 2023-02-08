package main

import "fmt"

func main() {
	//clave valor, las claves son unicas claramente
	ma1 := make(map[int]string)
	ma1[1] = "Hey"
	ma1[2] = "Brother"

	fmt.Println(ma1)
	fmt.Println(ma1[2])

	//borrar
	delete(ma1, 1)
	fmt.Println(ma1)
	ma1[2] = "Sister"
	fmt.Println(ma1)

	_, ok := ma1[100]
	fmt.Println("Existe :", ok)

	m2 := map[int]string{
		2: "A",
		3: "B",
		1: "C",
	}

	fmt.Println(m2)
}
