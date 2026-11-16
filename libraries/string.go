package libraries

func Pointer() {
	var umur int = 20
	println("value", umur)
	println("address", &umur)

	var umurAlias *int = &umur
	println("value", umurAlias)
	println("address", &umurAlias)
	println("value by pointer", *umurAlias)
}

func NamaFunction(nama  *string, age *int) {
	println(nama, age)
	println(*nama, *age)
}