package main

import (
	"fmt"
)

func main() {
	//Ciclo for es el unico que existe en Go
	sum := 0
	how := 0

	for i := 0; i < 10; i++ {
		sum = sum + i
		how++
		//fmt.Println(sum)
	}
	fmt.Println(how)
	fmt.Println(sum)
	/*Ciclo infinito
	for {
		fmt.Println("Sigue con juicio bro!")
	}*/

	/*en mi pc tardo 11:10:06
	counter := 0
	for {

		if counter >= 1000000 {
			break
		}
		counter++
		fmt.Println(counter)
	}
	*/

	numbers := []int{1, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	for i := range numbers {
		fmt.Println(numbers[i])
	}

	for i, v := range numbers {
		fmt.Println("Indice ", i, "valor ", v)
	}

	storee := make(map[string]int)
	storee["phones"] = 10
	storee["mouses"] = 25

	for i, v := range storee {
		fmt.Println("clave: ", i, "Valor ", v)
	}

	array := [5]int{4, 2, 5, 6, 7}
	for i, v := range array {
		array[i] = v + 20
	}
	fmt.Println("Respuesta tarea")
	fmt.Println(array)
	/*
		fmt.Println("Digite un numero")
		var eleccion int //Declarar variable y tipo antes de escanear, esto es obligatorio
		fmt.Scanln(&eleccion)*/
	//fmt.Println("Eleccion : ", eleccion)

	var slice_numbers []int
	for {
		fmt.Println("Digite un numero")
		var eleccion int //Declarar variable y tipo antes de escanear, esto es obligatorio
		fmt.Scanln(&eleccion)
		if eleccion == 0 {
			break
		}
		fmt.Println("Por favor digite 0 para detener el programa")
		slice_numbers = append(slice_numbers, eleccion)
	}
	fmt.Println("Slice numbers: ", slice_numbers)

	//tarea 3

	elements := map[string]string{
		"10": "notebook",
		"15": "tv",
		"21": "heladera",
		"27": "monitor",
		"35": "camara",
	}
	var respuestas []string
	//fmt.Println(elements)
	for {
		fmt.Println("Digite un numero")
		var eleccion string
		fmt.Scanln(&eleccion)
		if eleccion == "0" {
			break
		}

		_, ok := elements[eleccion]
		fmt.Println("Existe :", ok)
		fmt.Println(elements[eleccion])

		if ok == true {
			//var element string = elements[eleccion]
			respuestas = append(respuestas, string(elements[eleccion]))
		} else {
			respuestas = append(respuestas, "No encontrado")
		}
	}
	fmt.Println("Codigos por consola")
	fmt.Println(respuestas)

}
