package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {
	//return false // TODO: replace this
	var total1, total2 int
	//	data := make(map[string]int)

	total1 = 4*s[i].Correct - 1*s[i].Wrong
	total2 = 4*s[j].Correct - 1*s[j].Wrong
	/*
	   	- Jika ada yang nilainya sama, maka:
	     	- Yang `Jumlah Benar`-nya lebih tinggi akan diurutkan di atas.
	     	- Jika masih sama:
	         	- Yang `Nama`-nya lebih awal akan diurutkan di atas
	*/
	if total1 == total2 && s[i].Correct > s[j].Correct {
		return s[i].Name > s[j].Name
	} else if total1 == total2 && s[i].Correct == s[j].Correct {
		return s[i].Name < s[j].Name
	} else {
		return total1 > total2
	}
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	sort.Sort(s)
	names := []string{}

	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},
		{"Agus", 3, 4, 0},
		{"Doni", 3, 0, 7},
		{"Ega", 3, 0, 7},
		{"Anton", 2, 0, 5},
	})

	sort.Sort(scores)
	fmt.Println(scores.TopStudents())
}
