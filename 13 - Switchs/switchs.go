package main

import (
	"fmt"
)

func WeekDay(day int8) string {

	// Break clausule not exists in Golang
	switch day {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid weekday provided!"
	}
}

func main() {
	fmt.Println("Switchs")

	weekDay := WeekDay(7)

	fmt.Println("Today is", weekDay)
}
