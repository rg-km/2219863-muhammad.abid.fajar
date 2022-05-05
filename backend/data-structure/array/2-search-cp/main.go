package main

import "fmt"

func main() {
	var cars1 = []string{"Toyota", "Honda", "Nissan", "BMW", "Chevy", "Ford"}
	var cars2 = []string{"Ford", "BMW", "Audi", "Mercedes"}

	result, err := SearchMatch(cars1, cars2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Matched car brand: ", result)
	}
}

func SearchMatch(arr1 []string, arr2 []string) ([]string, error) {
<<<<<<< HEAD
	var dummy []string
	for _, v := range arr1 {
		for _, l := range arr2 {
			if v == l {
				dummy = append(dummy, v)
			}
		}
	}
	if dummy == nil {
		return nil, fmt.Errorf("no match")
	}
	return dummy, nil // TODO: replace this
=======
	return []string{}, nil // TODO: replace this
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
