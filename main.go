package main

import (
	"fmt"
	"golang-latihan/libraries"
	"runtime"
	"time"
)

type Hitung func(a int, b int) int

func Sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	println(total)
}

func Switch() {
	fmt.Print("Go berjalan pada ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Selamat pagi!")
	case t.Hour() < 17:
		fmt.Println("Selamat sore.")
	default:
		fmt.Println("Selamat malam.")
	}
}

func main() {
	//pointer
	libraries.Pointer()

	//function
	name := "elfaro"
	age := 50

	libraries.NamaFunction(&name, &age)
	println(name, age)

	//lambda
	func(name string, age int) {
		println("Hello word", name, age)
	}("jack", 20)

	//closure
	var hitung = func(a int, b int) int {
		return a + b
	}

	jumlah := hitung(10, 20)
	println(jumlah)

	//type funct
	var myHitung Hitung
	myHitung = func(a int, b int) int {
		return a + b
	}
	jumlah2 := myHitung(55, 20)
	println(jumlah2)

	Sum(1)
	Sum(1, 2)
	Sum(2, 3, 4)

	nums := []int{1, 2, 3, 4}
	Sum(nums...)

	Switch()
}
