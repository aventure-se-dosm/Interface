package main

import (
	"fmt"

	"figuras/Formas"
)

func main() {
	circulo1 := Formas.Circulo{
		Raio: 1.0,
	}

	retangulo1 := Formas.Retangulo{
		Base:   2.0,
		Altura: 3.0,
	}

	//...

	fmt.Printf("%0.2f\n", circulo1.Area())
	fmt.Printf("%0.2f\n", retangulo1.Area())
}
