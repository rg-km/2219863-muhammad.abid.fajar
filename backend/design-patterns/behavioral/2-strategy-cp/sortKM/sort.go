package sortKM

//kali ini kita coba menggunakan strategy pattern untuk menentukan sort
//ada dua sort yaitu ascending dan descending sort

type Strategy interface {
	Sort([]int)
}

type SortKM struct {
	Strategy Strategy
}

func (s *SortKM) Sort(array []int) {
	// TODO: answer here
	asc := SortKMasc{&AscendingSort{}}
	des := SortKMdes{&DescendingSort{}}
	if asc.Strategy == s.Strategy {
		asc.Strategy.Sort(array)
		return
	} else if des.Strategy == s.Strategy {
		des.Strategy.Sort(array)
		return
	}
}

func (s *SortKM) SetStrategy(Strategy Strategy) {
	s.Strategy = Strategy
}
