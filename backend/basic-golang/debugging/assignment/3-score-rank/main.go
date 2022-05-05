package main

import "fmt"

func main() {
	/*
		Score Rank:
		90-100: A
		80-89: B
		70-79: C
		60-69: D
		0-59: E
	*/
	res := ScoreRank(50)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := ScoreRankCorrect(arr)
	// fmt.Println(resCorrect)
}

func ScoreRank(score int) string {
	if score > 100 || score < 0 {
		return "INVALID"
	}
	var rank string
	switch {
	case score > 60:
		rank = "D"
	case score > 70:
		rank = "C"
	case score > 80:
		rank = "B"
	case score > 90:
		rank = "A"
	default:
		rank = "E"
	}

	return rank
}

func ScoreRankCorrect(score int) string {
<<<<<<< HEAD
	if score > 100 || score < 0 {
		return "INVALID"
	}
	var rank string
	if score >= 90 && score <= 100 {
		rank = "A"
	} else if score >= 80 && score <= 89 {
		rank = "B"
	} else if score >= 70 && score <= 79 {
		rank = "C"
	} else if score >= 60 && score <= 69 {
		rank = "D"
	} else {
		rank = "E"
	}

	return rank
=======
	return "" // TODO: replace this
>>>>>>> 0a32055256f6fde63d12cce9d6bf4e9ec0eccbd2
}
