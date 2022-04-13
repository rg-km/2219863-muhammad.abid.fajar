package main

import (
	"errors"
	"fmt"
)

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi pop.

var ErrStackUnderflow = errors.New("stack underflow")

type Stack struct {
	// TODO: answer here
	Top  int
	Size int
	Data []int
}

func NewStack(size int) Stack {
	// TODO: answer here
	return Stack{
		Top:  -1,
		Size: size,
		Data: []int{},
	}
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
	if len(s.Data) == s.Size {
		return fmt.Errorf("error")
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
		return nil
	}
}

func (s *Stack) Pop() (int, error) {
	// TODO: answer here
	if len(s.Data) == 0 {
		return 0, ErrStackUnderflow
	} else {
		poppedValue := s.Data[s.Top]
		s.Top -= 1
		s.Data = s.Data[:len(s.Data)-1]
		return poppedValue, nil
	}
}
