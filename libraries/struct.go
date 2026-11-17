package libraries

import "fmt"

type Mahasiswa struct {
	Nama string
	Umur int
}

func Struct() {
	m := Mahasiswa{Nama: "elfaro", Umur: 20}
	fmt.Println("nama:", m.Nama, "umur:", m.Umur)

	m.Umur = 21
	fmt.Printf("data lengkap: %+v\n", m)
}
