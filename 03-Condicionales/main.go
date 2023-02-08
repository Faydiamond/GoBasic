package main

import "fmt"

func ret_bool() bool {
	return false
}

func main() {
	fmt.Println("Condicionales")
	Myage := 30
	if Myage > 30 {
		fmt.Println("Sos menor de 30 anos, tienes mucho por hacer.")
	} else {
		fmt.Println("Tenes 30 anos, vamos a trabajar fuerte")
	}

	if value := ret_bool(); value == true {
		fmt.Println("Es verdadera la variable")
	} else {
		fmt.Println("Es falsa la variable")
	}

	day := 5
	switch day {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	}

	number := 0

	if number < 0 {
		fmt.Println("Es un numero negativo")
	} else if number == 0 {
		fmt.Println("El numero digitado es 0")
	} else if number > 0 {
		fmt.Println("El numero es positivo")
	} else {
		fmt.Println("Esto no es un numero!")
	}

	license := true
	age := 15

	if license == true && age > 15 {
		fmt.Println("Puede seguir avanzando")
	} else if license == false && age >= 0 {
		fmt.Println("No puede seguir circulando")
	} else if license == true && age <= 15 {
		fmt.Println("No puede seguir circulando")
	}
}
