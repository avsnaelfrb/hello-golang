package libraries

import (
	"fmt"
	"time"
)

func Select() {
	cepat := make(chan string)
	lambat := make(chan string)

	go func() {
		cepat <- "data cepat"
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		lambat <- "data lambat"
	}()

	select {
	case msg := <-cepat:
		fmt.Println("diterima:", msg)
	case msg := <-lambat:
		fmt.Println("diterima:", msg)
	}
}
