package coffe

type Mocha struct {
	// TODO: answer here
	Coffe Coffe
}

func (m Mocha) GetCost() float64 {
	cost := m.Coffe.GetCost()
	// return cost + 0.20
	return cost + 1.00
	// return 0 // TODO: replace this
}

func (m Mocha) GetDescription() string {
	description := m.Coffe.GetDescription()
	return description + ", Mocha"
	// return "" // TODO: replace this
}

type Whipcream struct {
	// TODO: answer here
	Coffe Coffe
}

func (w Whipcream) GetCost() float64 {
	cost := w.Coffe.GetCost()
	return cost + 0.20
	// return 0 // TODO: replace this
}

func (w Whipcream) GetDescription() string {
	description := w.Coffe.GetDescription()
	// return "" // TODO: replace this
	return description + ", Whipcream"
}
