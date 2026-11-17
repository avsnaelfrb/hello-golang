package libraries

import "fmt"

type Bangun interface {
	Luas() float64
}

type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

func Interface() {
	var b Bangun = Persegi{Sisi: 4}
	fmt.Println("luas persegi:", b.Luas())
}
