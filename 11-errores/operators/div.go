package operators

func Div(x, y float64) float64 {
	if y == 0 {
		panic("No puedo dividir entre 0")
	}
	return x / y
}
