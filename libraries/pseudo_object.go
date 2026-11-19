package libraries

import "fmt"

type Akun struct {
	Nama string
	Umur int
}

func NewAkun(nama string, umur int) Akun {
	return Akun{Nama: nama, Umur: umur}
}

func (a Akun) Sapa() string {
	return fmt.Sprintf("Halo, saya %s (%d tahun)", a.Nama, a.Umur)
}

func PseudoObject() {
	a := NewAkun("elfaro", 20)
	fmt.Println(a.Sapa())
}
