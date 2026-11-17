package libraries

// Closure menyimpan fungsi sebagai nilai variabel.
func Closure() {
	var hitung = func(a int, b int) int {
		return a + b
	}

	jumlah := hitung(10, 20)
	println(jumlah)
}
