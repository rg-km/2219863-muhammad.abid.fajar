package main

import (
	"errors"
	"fmt"
)

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi peek.

var ErrEmptyStack = errors.New("stack is empty")

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

func (s *Stack) Peek() (int, error) {
	// TODO: answer here
	if s.Top == -1 {
		return 0, ErrEmptyStack
	} else {
		return s.Data[s.Top], nil
	}
}
