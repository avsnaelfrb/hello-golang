package libraries

import "fmt"

func Channel() {
	pesan := make(chan string)

	go func() {
		pesan <- "halo dari goroutine"
	}()

	fmt.Println("diterima:", <-pesan)
}
