package perusahaan

type VP struct {
	Subordinate []Employee
}

func (vp VP) GetSalary() int {
	return 20
}

func (vp VP) TotalDivisonSalary() int {
	var total int

	for _, e := range vp.Subordinate {
		total += e.GetSalary()
	}
	return total  + vp.GetSalary()// TODO: replace this
}
