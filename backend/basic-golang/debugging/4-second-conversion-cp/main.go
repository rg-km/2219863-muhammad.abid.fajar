package main

import "fmt"

func main() {
	/*
		Convert the given second to 00:00:00 hour minute second format

		Example Input/Output
		30 -> 00:00:30
		70 -> 00:01:10
		67812 -> 18:50:12
		678120 -> 188:22:00

	*/
	res := ConvertSecondToTimeString(30)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := ConvertSecondToTimeStringCorrect(800)
	// fmt.Println(resCorrect)
}

func ConvertSecondToTimeString(second int) string {
	hours := second / 3600
	minutes := second / 60

	timeString := fmt.Sprintf("%d:%d:%d", hours, minutes, second)
	return timeString
}

func ConvertSecondToTimeStringCorrect(second int) string {
	// return "" // TODO: replace this

	hours := second / 3600
	minutes := (second - (hours * 3600)) / 60
	second = second - (hours * 3600) - (minutes * 60)
	if hours < 10 && minutes < 10 && second < 10 {
		timeString := fmt.Sprintf("0%d:0%d:0%d", hours, minutes, second)
		return timeString
	} else if hours < 10 && minutes < 10 {
		timeString := fmt.Sprintf("0%d:0%d:%d", hours, minutes, second)
		return timeString
	} else if hours < 10 {
		timeString := fmt.Sprintf("0%d:%d:%d", hours, minutes, second)
		return timeString
	} else if minutes < 10 {
		timeString := fmt.Sprintf("%d:0%d:%d", hours, minutes, second)
		return timeString
	} else if second < 10 {
		timeString := fmt.Sprintf("%d:%d:0%d", hours, minutes, second)
		return timeString
	} else {
		timeString := fmt.Sprintf("%d:%d:%d", hours, minutes, second)
		return timeString
	}
}
