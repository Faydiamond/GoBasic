package main

import (
	"curso_golang/Faydiamond/introduction/08-functions/functions"
	"fmt"
)

func main() {

	fmt.Println(functions.Add(100, 200))
	functions.RepeatStrings(5, "Haro")

	res, err := functions.Calcul(functions.SUM, 5, 0)
	if err != nil {
		fmt.Println(err.Error()) //retorna el error en string
	} else {
		fmt.Printf("El resultado es %f  \n", res)
	}

	fmt.Printf("Sumatoria es igual a %f \n", functions.Pluss(5, 10, 100, 200, 500, 5, 100))

	fn := functions.FactoryOperation(functions.SUM)
	v := fn(2, 3)
	fmt.Println("Suma", v)

}
