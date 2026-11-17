package libraries

func Sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	println(total)
}

// Variadic memanggil Sum dengan jumlah argumen berbeda.
func Variadic() {
	Sum(1)
	Sum(1, 2)
	Sum(2, 3, 4)

	nums := []int{1, 2, 3, 4}
	Sum(nums...)
}
