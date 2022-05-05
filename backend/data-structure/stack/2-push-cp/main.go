package main

import (
	"errors"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

var ErrStackOverflow = errors.New("stack overflow")

type Stack struct {
	// TODO: answer here
<<<<<<< HEAD
	Top  int
	Size int
	Data []int
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func NewStack(size int) Stack {
	// TODO: answer here
<<<<<<< HEAD
	return Stack{
		Top:  -1,
		Size: size,
		Data: []int{},
	}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
<<<<<<< HEAD
	if len(s.Data) == s.Size {
		return ErrStackOverflow
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
		return nil
	}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
