package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Hola mundo!")

	//log.Panic("Error panic")
	//log.Fatal("My error")
	file, err := os.Create("mylog.txt")
	if err != nil {
		panic(err)
	}

	l := log.New(file, "Logger:", log.LstdFlags|log.Lshortfile) //file,prefijo , bandera
	l.Println("Test #1")
	l.Println("Test #2")
	l.Println("Test #3")
}
