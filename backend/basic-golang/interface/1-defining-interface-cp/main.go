package main

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk membuat interface dan mengimplementasikan interface.
// Buatlah interface Employee yang memiliki method signature GetBonus() int
// Buatlah implementasi interface Employee pada objek Manager, SeniorEngineer, dan JuniorEngineer.
// Pada objek Manager, SeniorEngineer, dan JuniorEngineer memiliki satu atribut yaitu BaseSalary.
// Untuk menghitung bonus terdapat tiga aturan sebagai berikut:
// Bonus untuk Manager adalah 3 * BaseSalary
// Bonus untuk SeniorEngineer adalah 2 * BaseSalary
// Bonus untuk JuniorEngineer adalah 1 * BaseSalary

// TODO: answer here

type EmployeeSearcher interface {
	FindByName(name string) (Employee, error)
	FindByOrigin(BaseSalary int) ([]Employee, error)
}

type Employee struct {
	Name       string
	BaseSalary int
}

type EmployeeSearcherWithMap struct {
	data map[int]Employee
}

func TotalEmployeeBonus(employees []Employee) int {
	// Hitunglah total bonus yang dikeluarkan dari list of Employee
	// TODO: answer here
	return 0
}

func main() {
	// Buatlah objek konkret untuk masing-masing objek dan panggil function TotalEmployeeBonus. Print total bonus untuk semua employee.
	// TODO: answer here
}
