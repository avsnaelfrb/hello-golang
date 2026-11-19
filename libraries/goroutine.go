package libraries

import (
	"fmt"
	"sync"
)

func Goroutine() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("goroutine", n)
		}(i)
	}
	wg.Wait()
	fmt.Println("semua goroutine selesai")
}
