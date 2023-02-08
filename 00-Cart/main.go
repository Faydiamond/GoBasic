package main

import (
	"curso_golang/Faydiamond/introduction/00-Cart/structs"
	"encoding/json"
	"fmt"
)

func main() {
	camera_vision := structs.Product{
		Id:      1,
		Name:    "Camera Hik Vision",
		Counter: 15,
		Price:   75455,
	}

	cart := structs.Cart{
		IdUser: 1,
	}
	cart.AddCart(camera_vision)

	fmt.Println("Carrito de compras")
	fmt.Println("Cuantos productos: ", len(cart.Products))
	fmt.Printf("Total : %f", cart.Total())

	jsonproduct, _ := json.Marshal(camera_vision)
	fmt.Println(string(jsonproduct))

}
