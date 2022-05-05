package main

<<<<<<< HEAD
import (
	"errors"
	"fmt"
)
=======
import "errors"
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi peek.

var ErrEmptyStack = errors.New("stack is empty")

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
		return fmt.Errorf("error")
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
		return nil
	}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func (s *Stack) Peek() (int, error) {
	// TODO: answer here
<<<<<<< HEAD
	if s.Top == -1 {
		return 0, ErrEmptyStack
	} else {
		return s.Data[s.Top], nil
	}
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
