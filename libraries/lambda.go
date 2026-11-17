package libraries

// Lambda menjalankan fungsi anonim yang langsung dipanggil.
func Lambda() {
	func(name string, age int) {
		println("Hello word", name, age)
	}("jack", 20)
}
