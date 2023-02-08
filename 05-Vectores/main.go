package main

import "fmt"

func main() {
	names := [3]string{"Fabian", "Daniel", "Juan"}
	fmt.Println(names)
	fmt.Println("Tamano array", len(names))
	var my_slice []int
	fmt.Println("mi slice ", my_slice)
	my_slice = append(my_slice, 5, 10, 15, 20, 25)
	fmt.Println("mi slice ", my_slice)
	//agarrar de un arreglo
	splitNames := names[0:2]
	fmt.Println("Split names ", splitNames, "Tamano ", len(splitNames))
	slice2 := make([]string, 3)
	fmt.Printf("Tamano %d , valor %v ", len(slice2), slice2)
}
