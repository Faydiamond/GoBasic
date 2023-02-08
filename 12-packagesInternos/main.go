package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	//acceder variables de entorno
	env1 := os.Getenv("USERNAME")
	p("Enviroment : ", env1)

	dir, _ := os.Getwd()
	fmt.Println("Direccion : ", dir)

}
