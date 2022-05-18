package sortKM

import "sort"

// TODO: answer here
type SortKMasc struct {
	Strategy Strategy
}

//concrete strategy implementation
type AscendingSort struct{}

func (as *AscendingSort) Sort(array []int) {
	//choose any sort algo you want
	// TODO: answer here
	sort.Sort(sort.IntSlice(array))
}
