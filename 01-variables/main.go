package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var number int
	number = -12
	fmt.Println(number)

	var numer_plus uint
	numer_plus = 12
	fmt.Println("Uint :", numer_plus)

	var cadena string
	cadena = "Mensaje"
	fmt.Println("cadena ", cadena)

	var flag bool
	flag = true
	fmt.Println("Boolean: ", flag)

	fmt.Println("Address of variable cadena ", &cadena)
	my_name := "Fabian"

	fmt.Println("My name is : ", my_name)

	const pi = 75.14
	fmt.Println("Constante ", pi)

	age := 29
	fmt.Printf("My age is: %d \n", age)
	fmt.Printf("My age is: %T, bytes:%d , bits: %d \n", age, unsafe.Sizeof(age), unsafe.Sizeof(age)*8)
	//solo viven en este ambito
	{

		floatvar := 3.24554

		fmt.Println("Que ve ", floatvar)
		strt := fmt.Sprintf("%f \n", floatvar)
		fmt.Printf("Tipo %T, valor: %s \n", strt, strt)
		//eliminar decimales
		strt2 := fmt.Sprintf("%.2f \n", floatvar)
		fmt.Println("Elimine dos decimales de ", floatvar, "Resultado: ", strt2)

		intvalue, _ := strconv.ParseInt("125", 0, 64)
		fmt.Printf("%d  %T \n", intvalue, intvalue)
	}
}
