package Formas

type Retangulo struct {
	Altura float64
	Base   float64
}

func (r Retangulo) Area() float64 {
	return r.Base * r.Altura
}
