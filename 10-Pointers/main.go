package main

import "fmt"

type hero struct {
	name     string
	age      uint
	typeee   string
	isactive bool
}

func main() {
	var i int
	var ip *int //puntero
	i = 34
	ip = &i
	fmt.Println(&i) //referencia -> direccion de memoria
	fmt.Println(ip)
	fmt.Println(i)
	fmt.Println(i) //referencia -> direccion de memoria
	fmt.Println(*ip)
	//%p

	fmt.Println("Puntero")
	thor := &hero{name: "thor", age: 29, isactive: true, typeee: "Electrico"}
	fmt.Printf("addrs: variable thor %p\n", thor)
	thor.set_hero()
	thor.set_hero_pointer()

}

func (h hero) set_hero() {
	fmt.Printf("addrs: %p normal \n", &h)
	h.age = 29
	h.name = "Thor"
}

func (h *hero) set_hero_pointer() {
	fmt.Printf("addrs: %p  with pointer \n", h)
	h.age = 29
}
