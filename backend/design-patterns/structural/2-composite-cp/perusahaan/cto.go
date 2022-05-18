package perusahaan

type CTO struct {
	Subordinate []Employee
}

func (c CTO) GetSalary() int {
	return 30
}

func (c CTO) TotalDivisonSalary() int {
	var total int

	for _, k := range c.Subordinate {
		total += k.TotalDivisonSalary()
	}

	return total + c.GetSalary()// TODO: replace this
}
