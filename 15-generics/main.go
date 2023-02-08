package main

import "fmt"

type CustomSlice[V int | string] []V

type Point int

//con el gorrito se me acepta cuando le envio alguno de estas opcones
type N1 interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func plus(num ...int) int {
	var res int
	for i := 0; i < len(num); i++ {
		res += num[i]
	}
	return res
}

func MinNumber[R N1](x, y R) R {
	if x < y {
		return x
	}
	return y
}

//Generics
func pluss[M int | int8 | int16 | int32 | int64 | float32 | float64](a []M) M {

	var total M
	for _, v := range a {
		total += v
	}
	return total
}

func pluss2[M Number](a []M) M {

	var total M
	for _, v := range a {
		total += v
	}
	return total
}

func anytype[N any](a, b N) {
	fmt.Println(a, " ", b)
}

//con los any no se puede comparar
func anytype2[N1 any, N2 any](a N1, b N2) {
	fmt.Println(a, " ", b)
}

//solo se pueden comprar  igual y distinto con el comparable
func comparablee[N comparable](a, b N) {
	//fmt.Println(a, " ", b)
	fmt.Println(a == b)
}

func main() {
	numbers := []int{2, 5, 5, 5, 3, 50, 100}
	numbers2 := []float32{2.5, 5.3, 5.1, 5.6, 4.7, 50.2, 100.6}
	fmt.Println(plus(numbers...))
	fmt.Println()
	fmt.Println(pluss(numbers))
	fmt.Println()
	fmt.Println(pluss(numbers2))
	fmt.Println("Otra manera")
	fmt.Println(pluss2(numbers2))

	//generico, no puedo enviar valores de distintos tipos, me toma el primero que envio como generico
	anytype(5, 5)
	anytype2(5, "Ohhh")
	comparablee(5, 5)

	//vectorr custom
	mySliceInt := CustomSlice[int]{1, 2, 3}
	fmt.Println(mySliceInt)
	mySliceStr := CustomSlice[string]{"a", "b", "c"}
	fmt.Println(mySliceStr)

	//enviar un interface
	c, y := Point(50), Point(10)
	fmt.Println(MinNumber(c, y))
}
