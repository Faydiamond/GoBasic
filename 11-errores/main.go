package main

import (
	"curso_golang/Faydiamond/introduction/11-errores/operators"
	"errors"
	"fmt"
)

func main() {
	var err error

	fmt.Println(err)
	err = errors.New("Mi error")
	fmt.Println(err)
	fmt.Println(err.Error())

	err_twoo := fmt.Errorf("Error %s  valor %d ", "x", 000)
	fmt.Println(err_twoo.Error())

	//funcion anonima
	defer func() { //defer se ejecuta siempre al; finalizar la funcion
		fmt.Println("defender")
		//manejo el error
		cov := recover()
		if cov != nil {
			fmt.Println("Recovered in \n ", cov)
		}
	}()
	x := 12.0
	j := -2.0

	res := operators.Div(x, j)
	fmt.Println("resultado : ", res)

}
