package libraries

import "fmt"

type Kalkulator struct{}


func (Kalkulator) Tambah(angka ...int) int {
	total := 0
	for _, n := range angka {
		total += n
	}
	return total
}

func MethodOverloading() {
	k := Kalkulator{}
	fmt.Println("tambah 1 angka:", k.Tambah(5))
	fmt.Println("tambah 3 angka:", k.Tambah(5, 3, 2))
}
