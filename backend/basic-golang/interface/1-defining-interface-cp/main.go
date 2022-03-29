package main

import (
	"fmt"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk membuat interface dan mengimplementasikan interface.
// Buatlah interface Employee yang memiliki method signature GetBonus() int
// Buatlah implementasi interface Employee pada objek Manager, SeniorEngineer, dan JuniorEngineer.
// Pada objek Manager, SeniorEngineer, dan JuniorEngineer memiliki satu atribut yaitu BaseSalary.
// Untuk menghitung bonus terdapat tiga aturan sebagai berikut:
// Bonus untuk Manager adalah 3 * BaseSalary
// Bonus untuk SeniorEngineer adalah 2 * BaseSalary
// Bonus untuk JuniorEngineer adalah 1 * BaseSalary

// TODO: answer here

type Employee struct {
	Employee []Employee
}
type Manager struct {
	BaseSalary int
}
type SeniorEngineer struct {
	BaseSalary int
}
type JuniorEngineer struct {
	BaseSalary int
}

func main() {
	// Buatlah objek konkret untuk masing-masing objek dan panggil function TotalEmployeeBonus. Print total bonus untuk semua employee.
	// TODO: answer here

	manager := Manager{
		BaseSalary: 20000000,
	}
	bonus1 := manager.GetBonus()
	fmt.Println(bonus1)

	seniorEngineer := SeniorEngineer{
		BaseSalary: 15000000,
	}
	bonus2 := seniorEngineer.GetBonus()
	fmt.Println(bonus2)

	juniorEngineer := JuniorEngineer{
		BaseSalary: 10000000,
	}
	bonus3 := juniorEngineer.GetBonus()
	fmt.Println(bonus3)

	// g := bonus1 + bonus2 + bonus3
	// totalBonus := TotalEmployeeBonus(make([]Employee, g))
	// fmt.Println(totalBonus)

}

// func TotalEmployeeBonus(employee []Employee) int {
// 	// Hitunglah total bonus yang dikeluarkan dari list of Employee
// 	// TODO: answer here

// 	// Manager := 3 * 20_000_000
// 	// SeniorEngineer := 2 * 15_000_000
// 	// JuniorEngineer := 10_000_000
// 	// b := Manager + SeniorEngineer + JuniorEngineer
// 	// return b
// 	res := make([]Employee, len(employee))
// 	for _, x := range employee {
// 		res = append(res, x)
// 	}

// 	var j int
// 	return j
// }

func (manager Manager) GetBonus() int {
	return manager.BaseSalary * 3
}
func (sen SeniorEngineer) GetBonus() int {
	return sen.BaseSalary * 2
}
func (jun JuniorEngineer) GetBonus() int {
	return jun.BaseSalary
}
