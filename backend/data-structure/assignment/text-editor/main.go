package main

import (
<<<<<<< HEAD
=======
	"fmt"

>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	"github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"
)

// Implementasikan teks editor sederhana
// Teks editor ini memiliki 4 method
// Write: digunakan untuk menulis per karakter
// Read: digunakan untuk mencetak karakter yang telah ditulis
// Undo: digunakan untuk melakukan operasi undo
// Redo: digunakan untuk melakukan operasi redo

type TextEditor struct {
	UndoStack *stack.Stack
	RedoStack *stack.Stack
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		UndoStack: stack.NewStack(),
		RedoStack: stack.NewStack(),
	}
}

func (te *TextEditor) Write(ch rune) {
	// TODO: answer here
<<<<<<< HEAD
	te.UndoStack.Push(ch)

=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func (te *TextEditor) Read() []rune {
	// TODO: answer here
<<<<<<< HEAD
	var r []rune
	r = append(r, te.UndoStack.Data...)
	return r
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func (te *TextEditor) Undo() {
	// TODO: answer here
<<<<<<< HEAD
	checkundo, err := te.UndoStack.Peek()
	if err != nil {
		return
	}
	te.UndoStack.Pop()
	te.RedoStack.Push(checkundo)
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}

func (te *TextEditor) Redo() {
	// TODO: answer here
<<<<<<< HEAD
	checkredo, err := te.RedoStack.Peek()
	if err != nil {
		return
	}
	te.UndoStack.Push(checkredo)
	te.RedoStack.Pop()
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
