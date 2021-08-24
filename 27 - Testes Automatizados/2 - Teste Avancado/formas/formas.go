package formas

import (
	"fmt"
	"math"
)

//Retangulo - struct de retangulo
type Retangulo struct {
	Altura  float64
	Largura float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

//Circulo - struct de circulo
type Circulo struct {
	raio float64
}

type Forma interface {
	area() float64
}
