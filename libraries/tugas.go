package libraries

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func firstPrimes(n int, start ...int) []int {
	dari := 2
	if len(start) > 0 {
		dari = start[0]
	}

	prima := make([]int, 0, n)
	for angka := dari; len(prima) < n; angka++ {
		if isPrime(angka) {
			prima = append(prima, angka)
		}
	}
	return prima
}

func each(nums []int, visit func(int)) {
	for _, n := range nums {
		visit(n)
	}
}

func Tugas(n int, count *int) {
	prima := firstPrimes(n)

	label := func(p int) string {
		switch {
		case p < 10:
			return "satuan"
		case p < 100:
			return "puluhan"
		default:
			return "ratusan"
		}
	}

	printed := 0
	each(prima, func(p int) {
		if p > 0 {
			fmt.Printf("%d (%s) ", p, label(p))
			printed++
		}
	})
	fmt.Println()

	*count = printed
}
