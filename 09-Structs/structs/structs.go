package structs

type Product struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Type Type   `json:"type"`

	Tags  []string `json:"tags"`
	Price float32  `json:"price"`
	Count int      `json:"count"`
}

type Type struct {
	Id          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Car struct {
	IdUser   uint      `json:"iduser"`
	Id       uint      `json:"id"`
	Products []Product `json:"productos"`
}

func (p Product) TotalPrice() float64 {
	return float64(p.Count) * float64(p.Price)
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) AddTaggs(tag ...string) {
	p.Tags = append(p.Tags, tag...)
}

func (c *Car) AddCar(products ...Product) {
	c.Products = append(c.Products, products...)
}

func (c Car) PriceAll() float64 {
	var total float64
	for _, v := range c.Products {
		total += v.TotalPrice()
	}
	return total
}

func NewCart(iduser uint) Car {
	return Car{IdUser: iduser}
}
