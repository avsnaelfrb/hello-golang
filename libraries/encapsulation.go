package libraries

import (
	"errors"
	"fmt"
)

type Rekening struct {
	saldo int
}

func (r *Rekening) Setor(jumlah int) error {
	if jumlah <= 0 {
		return errors.New("jumlah setor harus lebih dari 0")
	}
	r.saldo += jumlah
	return nil
}

func (r *Rekening) Saldo() int {
	return r.saldo
}

func Encapsulation() {
	var r Rekening
	if err := r.Setor(50000); err != nil {
		fmt.Println("gagal:", err)
	}
	if err := r.Setor(-1000); err != nil {
		fmt.Println("gagal:", err)
	}
	fmt.Println("saldo:", r.Saldo())
}
