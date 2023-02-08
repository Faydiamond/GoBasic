package main

import (
	"fmt"
	"os"
)

func Open_file() {
	file, err := os.Open("filee.txt")
	if err != nil {
		fmt.Println("Error 1")
		os.Exit(1) //error
	}
	data := make([]byte, 200)

	c, err := file.Read(data)
	if err != nil {
		fmt.Println("Error 2")
		os.Exit(1)
	}
	fmt.Printf("%q ", data[:c])
}

func main() {
	Open_file()
}
