package structs

type Product struct {
	Id      uint    `json:"id"`
	Name    string  `json:"name"`
	Counter int     `json:"counter"`
	Price   float64 `json:"price"`
}

type Cart struct {
	IdUser   uint
	Products []Product
}

func (c *Cart) AddCart(product ...Product) {
	c.Products = append(c.Products, product...)
}

func (c *Cart) Total() float64 {
	var total float64
	for _, v := range c.Products {
		total += v.TotalPrice()
	}
	return total
}

func (p Product) TotalPrice() float64 {
	return float64(p.Price * float64(p.Counter))
}
