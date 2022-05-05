// Golang program to count digits of given number
// using recursion

package main

import "fmt"

var count int = 0

//function to count digits
func CountDigits(num int) int {

	if num > 0 {
<<<<<<< HEAD
		// TODO: replace this
		count++
		CountDigits(num / 10)
=======
		CountDigits(0) // TODO: replace this
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
	}
	return count

}

func main() {
	var num int = 0
	var result int = 0

	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)

	result = CountDigits(num)
	fmt.Printf("Count of digits is: %d\n", result)
}
