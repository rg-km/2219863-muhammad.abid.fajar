package main

import (
	"errors"
)

// Penambahan elemen pada stack disebut dengan push.
// Ketika menambah elemen, nilai pada top bertambah 1.

type Stack struct {
	Top  int
	Data []int
}

func (s *Stack) Push(Elemen int) {
	s.Top += 1
	s.Data = append(s.Data, Elemen)
}

// Jika menggunakan array yang memiliki jumlah maksimal elemen, maka perlu melakukan pengecekan apakah elemen sudah penuh atau belum.
type StackWithArray struct {
	Top  int
	Data [5]int
}

func (s *StackWithArray) Push(Elemen int) error {
<<<<<<< HEAD
	if s.Top == len(s.Data) {
=======
	if s.Top == len(s.Data)-1 {
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
		return errors.New("stack overflow")
	} else {
		s.Top += 1
		s.Data[s.Top] = Elemen
	}
	return nil
}
