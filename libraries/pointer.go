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
