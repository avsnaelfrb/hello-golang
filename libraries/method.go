package libraries

import (
	"fmt"
	"math"
)

type Lingkaran struct {
	JariJari float64
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * l.JariJari * l.JariJari
}

func (l *Lingkaran) Perbesar(kali float64) {
	l.JariJari *= kali
}

func Method() {
	l := Lingkaran{JariJari: 7}
	fmt.Println("luas:", l.Luas())

	l.Perbesar(2)
	fmt.Println("jari-jari setelah diperbesar:", l.JariJari)
}
