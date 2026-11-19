package libraries

import (
	"fmt"
	"sync"
)

// Goroutine menjalankan 3 pekerjaan secara bersamaan. WaitGroup dipakai
// supaya main menunggu semua goroutine selesai sebelum program berhenti.
func Goroutine() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("goroutine", n)
		}(i) // i dikirim sebagai parameter, bukan ditangkap langsung dari loop
	}
	wg.Wait()
	fmt.Println("semua goroutine selesai")
}
