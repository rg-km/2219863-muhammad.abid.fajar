package main

import (
	"errors"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

var ErrStackOverflow = errors.New("stack overflow")

type Stack struct {
	// TODO: answer here
	Top  int
	Data []int
}

func NewStack(size int) Stack {
	// TODO: answer here
	var s Stack
	s.Top += 1
	s.Data = append(s.Data, size)
	return s
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
	if s.Top == len(s.Data) {
		return errors.New("stack overflow")
	} else {
		s.Top += 1
		s.Data[s.Top] = Elemen
	}
	return nil
}
