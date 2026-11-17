package libraries

import "fmt"

func InterfaceNil() {
	var kosong Bangun
	fmt.Println("interface kosong == nil:", kosong == nil)

	// Pointer nil yang dimasukkan ke interface membuat interface ikut terisi,
	// jadi hasilnya false walaupun nilai di dalamnya nil.
	var p *Persegi
	var terisi Bangun = p
	fmt.Println("typed nil == nil:", terisi == nil)
}
