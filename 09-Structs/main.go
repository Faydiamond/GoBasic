package main

import (
	"curso_golang/Faydiamond/introduction/09-Structs/structs"
	"curso_golang/Faydiamond/introduction/09-Structs/vehicles"
	"encoding/json"
	"fmt"
)

func main() {
	Pc := structs.Product{
		Id:    1,
		Name:  "Portatil Hp",
		Type:  structs.Type{Id: 1, Code: "Rx01", Description: "Desc"},
		Tags:  []string{"Portatiles", "hp"},
		Price: 2255999,
		Count: 10,
	}

	ram := structs.Product{
		Id:    2,
		Name:  "ram ddr4",
		Type:  structs.Type{Id: 1, Code: "Rx02", Description: "Desc"},
		Tags:  []string{"Memorias ram", "ddr4"},
		Price: 150000,
		Count: 12,
	}
	fmt.Println(Pc)

	var wilson_bag structs.Product
	wilson_bag.Id = 2
	wilson_bag.Name = "Maleta Wilson"

	fmt.Println(wilson_bag)

	v, e := json.Marshal(Pc)
	fmt.Println(e)
	fmt.Println(string(v))

	fmt.Println(Pc.TotalPrice())
	Pc.SetName("Toshiba")
	fmt.Println(Pc)

	wilson_bag.AddTaggs("Mochilas", "Wilson", "Estudiantil", "Ejecutiva", "Viajero")

	fmt.Println(wilson_bag)

	car := structs.NewCart(5)
	car.AddCar(Pc, ram)
	fmt.Println("Carrito de compras")
	fmt.Println("Cantidad de productos: ", len(car.Products))
	fmt.Printf("Total : %f", car.PriceAll())

	fmt.Println("Vehiculos ///////////////////////////////////////////////////")
	mycar := vehicles.Auto{
		Time: 160,
	}
	var d float64
	fmt.Println(mycar.Distance())
	vehicless := []string{"Auto", "Avion", "Camion", "Tractor", "Tren", "Barco"}
	for _, v := range vehicless {
		fmt.Printf("Vehiculo:  %s  \n", v)
		vec, err := vehicles.SetVehicle(v, 400)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println()
			continue
		}
		distancee := vec.Distance()
		fmt.Printf("Distancia : %2.f \n", distancee)
		fmt.Println()
		d += distancee
	}

	fmt.Printf("Total distancia : %.2f \n ", d)
	fmt.Println()
}
