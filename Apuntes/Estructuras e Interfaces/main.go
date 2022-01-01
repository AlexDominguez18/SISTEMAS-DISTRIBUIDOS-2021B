package main

import (
	"fmt"
	"./figuras"
)

func sumAreas(figs ...figuras.Figura) float64 {
	areaTotal := 0.0

	for _, f := range figs {
		areaTotal += f.Area()
	}

	return areaTotal
}

func main() {
	c1 := figuras.Circulo{Radio: 5}
	r1 := figuras.Rectangulo{Base: 10, Altura: 5}
	// c2 := Circulo{radio: 10}
	// c3 := Circulo{15}
	// c4 := new(Circulo)
	// c5 := &Circulo{100}

	// fmt.Println(c1)
	// fmt.Println(c2)
	// fmt.Println(c3)
	// fmt.Println(c4)
	// fmt.Println(c5)

	fm := figuras.FiguraMultiple{
		Figuras: []figuras.Figura{
			&c1, 
			&r1, 
			&figuras.Circulo{Radio: 2},
		},
	}

	fmt.Println(c1.Area())
	fmt.Println(r1.Area())
	fmt.Println(sumAreas(&c1,&r1))
	fmt.Println(fm.Area())
}