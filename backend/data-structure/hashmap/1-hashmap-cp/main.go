package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{{name: "Bob", age: 21}, {name: "Sam", age: 28}, {name: "Ann", age: 21}, {name: "Joe", age: 22}, {name: "Ben", age: 28}}
	fmt.Println(AgeDistribution(people))
	fmt.Println(FilterByAge(people, 21))
}

func AgeDistribution(people []Person) map[int]int {
<<<<<<< HEAD
	dummy := map[int]int{
		22: 0,
		21: 0, 
		28: 0,
	}
	for _, v := range people {
		if v.age == 22 {
			dummy[v.age]++
		} else if v.age == 21 {
			dummy[v.age]++
		} else {
			dummy[v.age]++
		}
	}
	return dummy
}

func FilterByAge(people []Person, age int) []Person {
	var dummy []Person
	for _, v := range people {
		if v.age == age {
			dummy = append(dummy, v)
		}
	}

	return dummy
=======
	return map[int]int{} // TODO: replace this
}

func FilterByAge(people []Person, age int) []Person {
	return []Person{} // TODO: replace this
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
