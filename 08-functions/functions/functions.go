package functions

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(a, b int32) int32 {
	return a + b
}

func RepeatStrings(i int, msg string) {
	for x := 0; x < i; x++ {
		fmt.Println(msg)
	}
}

func Calcul(op Operation, x, y float64) (float64, error) {

	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case MUL:
		return x * y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("No se puede dividir entre cero")
		}
		return x / y, nil
	}
	return 0, errors.New("Se ha presentadop un erro")
}

func Pluss(values ...float64) float64 {
	var SUM float64
	for _, v := range values {
		SUM += v
	}
	return SUM
}
