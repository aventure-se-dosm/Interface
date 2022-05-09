package Formas

type Triangulo struct {
	Altura float64
	Base   float64
}

func (r Triangulo) Area() float64 {
	return float64(r.Base * r.Altura / 2.0)
}
