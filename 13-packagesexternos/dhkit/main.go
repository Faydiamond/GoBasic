package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lasstname"`
	Age       uint   `json:"age"`
	Mail      string `json:"mail"`
}

func main() {
	myUser := User{
		FirstName: "Daniel",
		LastName:  "Melero",
		Mail:      "DaniMelero@gmail.com",
		Age:       45,
	}

}

func print(v interface{}) {
	j, _ := json.Marshal(v)
	fmt.Println(string(j))
	fmt.Println()
}
