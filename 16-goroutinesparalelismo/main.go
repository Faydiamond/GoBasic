package main

import (
	"fmt"
	"time"
)

func main() {
	//trabajamos en paralelo proceso 1 y el 2 a la ves
	//tenemos un hilo por cada nucleo
	//canales enviar y recibir info mediante gorotines
	go myProcess("Prueba1")
	go myProcess("Prueba2")

	time.Sleep(10 * time.Second)

	myFirstChannel := make(chan string)  //declaro un canal
	myFirstChannel1 := make(chan string) //declaro un canal
	myFirstChannel2 := make(chan string) //declaro un canal

	go myProcessWithChannel("C", myFirstChannel)
	go myProcessWithChannel("D", myFirstChannel1)
	go myOtherWithChannel("E", myFirstChannel2)

	result := <-myFirstChannel //acabo el canal
	fmt.Println(result)
	close(myFirstChannel1)
	close(myFirstChannel2)
	close(myFirstChannel)
}

func myProcess(s string) {
	i := 0
	for i < 200 {
		time.Sleep(15 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d \n", s, i)
	}
}

func myProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 200 {
		time.Sleep(15 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d \n", p, i)
	}
	//cierro el canal
	c <- "ok"
}

func myOtherWithChannel(p string, c chan string) {
	i := 0
	for i < 100 {
		time.Sleep(25 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d \n", p, i)
	}
	//cierro el canal
	c <- "ok2"
}
