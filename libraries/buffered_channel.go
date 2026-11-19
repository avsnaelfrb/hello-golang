package libraries

import "fmt"

func BufferedChannel() {
	pesan := make(chan string, 2)

	pesan <- "satu"
	pesan <- "dua"

	fmt.Println("len:", len(pesan), "cap:", cap(pesan))
	fmt.Println("diterima:", <-pesan)
	fmt.Println("diterima:", <-pesan)
}
