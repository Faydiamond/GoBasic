package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name       string
	Age        uint
	Job        string
	Experience uint
}

func main() {
	myInt := 10
	myPnt := &myInt
	Examine(myInt)
	Examine(myPnt)

	myUser := User{Name: "Fabian", Age: 29, Job: "Developer", Experience: 3}
	Examine(myUser)
}

func Examine(data interface{}) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	k := t.Kind() //clase

	if reflect.ValueOf(data).Kind() == reflect.Struct {
		v := reflect.ValueOf(data)
		//fmt.Println("Number of fields: ", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			//fmt.Printf("Field:%d  type:%T value:%v  \n", i, v.Field(i), v.Field(i))

			switch v.Field(i).Kind() {
			case reflect.Int, reflect.Float32, reflect.Int64, reflect.Int8:
				myVar := v.Field(i).Int()
				fmt.Printf("Field:%d  type:%T value:%v  \n", i, myVar, myVar)
			case reflect.String:
				myVar := v.Field(i).String()
				fmt.Printf("Field:%d  type:%T value:%v  \n", i, myVar, myVar)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				myVar := v.Field(i).Uint()
				fmt.Printf("Field:%d  type:%T value:%v  \n", i, myVar, myVar)
			default:
				fmt.Println("Unsoported type ", v.Field(i).Type())
			}
		}
	} else {
		fmt.Println("Type : ", t)
		fmt.Println("Value : ", v)
		fmt.Println("Kind : ", k)
		fmt.Println()
	}
}
