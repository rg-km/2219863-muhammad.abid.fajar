// Reverse

package main

import "fmt"

func Reverse(st []string, depth int) string {
	str := ""
	// TODO: answer here
<<<<<<< HEAD
	if depth == -1 {
		return str
	}
	str += st[depth]
	str += Reverse(st, depth-1)
=======
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	return str
}

func main() {
	str := []string{"A", "I", "S", "E", "N", "O", "D", "N", "I"}
	s := Reverse(str, len(str)-1)
	fmt.Println(s)
}
