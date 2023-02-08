package vehicles

import "fmt"

type Vehicle interface {
	Distance() float64
}
type Auto struct {
	Time int
}

const (
	CarVehicle = "Auto"
	Plane      = "Avion"
	Truck      = "Camion"
	Tractor    = "Tractor"
	Train      = "Tren"
	Barco      = "Ship"
)

type Planee struct {
	Time int
}

type Truckk struct {
	Time int
}

type Trainn struct {
	Time int
}

type Tractorr struct {
	Time int
}

type Shipp struct {
	Time int
}

func (c *Auto) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

func (p *Planee) Distance() float64 {
	return 250 * (float64(p.Time) / 60)
}

func (t *Truckk) Distance() float64 {
	return 60 * (float64(t.Time) / 60)
}

func (tr *Tractorr) Distance() float64 {
	return 25 * (float64(tr.Time) / 60)
}

func (tn *Trainn) Distance() float64 {
	return 200 * (float64(tn.Time) / 60)
}

func (s *Shipp) Distance() float64 {
	return 80 * (float64(s.Time) / 60)
}

func SetVehicle(v string, time int) (Vehicle, error) {
	fmt.Println("Que mierda ve : ", v, time)

	switch v {
	case CarVehicle:
		return &Auto{Time: time}, nil
	}
	return nil, fmt.Errorf("vehiculo  %s  no existe", v)
}
