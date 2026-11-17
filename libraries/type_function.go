package libraries

// Hitung adalah tipe fungsi dengan dua parameter int dan satu hasil int.
type Hitung func(a int, b int) int

// TypeFunction memakai tipe Hitung sebagai tipe variabel.
func TypeFunction() {
	var myHitung Hitung
	myHitung = func(a int, b int) int {
		return a + b
	}
	jumlah2 := myHitung(55, 20)
	println(jumlah2)
}
