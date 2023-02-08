package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	//env the other file
	if err := godotenv.Load("myfolder/.var"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV4"))
	fmt.Println(os.Getenv("MY_ENV5"))
	fmt.Println(os.Getenv("MY_ENV6"))

	myEnv, err := godotenv.Read("myfolder/.var")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myEnv)
	//overload reemplaza enviroments
}
