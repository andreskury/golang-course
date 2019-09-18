package main

import (
	"fmt"
	"math"
)

type Figura interface {
	Area() float32
	Perimetro() float32
}

type Rectangulo struct {
	Ancho, Alto float32
}

func (r Rectangulo) Area() float32 {
	return r.Ancho * r.Alto
}

func (r Rectangulo) Perimetro() float32 {
	return 2*r.Ancho + 2*r.Alto
}

type Circulo struct {
	Radio float32
}

func (c Circulo) Area() float32 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float32 {
	return math.Pi * c.Radio * 2
}

func MostrarFigura(f Figura) {
	fmt.Printf("%#v - Area: %f, Perimetro: %f\n", f, f.Area(), f.Perimetro())
}

func main() {
	r := Rectangulo{Ancho: 3, Alto: 4}
	c := Circulo{Radio: 5}

	MostrarFigura(r)
	MostrarFigura(c)
}
