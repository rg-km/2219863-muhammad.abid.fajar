package sortKM

import "sort"

// TODO: answer here
type SortKMdes struct {
	Strategy Strategy
}

//concrete strategy implementation
type DescendingSort struct{}

func (ds *DescendingSort) Sort(array []int) {
	//choose any sort algo you want
	// TODO: answer here
	sort.Slice(array, func(i, j int) bool {
		return array[j] < array[i]
	})
}
