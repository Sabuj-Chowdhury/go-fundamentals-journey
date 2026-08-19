package main

import "fmt"

type WeekDays int

// int enum types
const (
	Monday WeekDays = iota
	Tuesday
	wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func getWorkingDay(day WeekDays) string {
	switch day {
	case Monday, Tuesday, wednesday, Thursday, Sunday:
		return "Office is open"
	case Friday, Saturday:
		return "Office is Closed"
	default:
		return "Invalid Day"

	}
}

type status string

const (
	Open   status = "open"
	Closed status = "closed"
)

func main() {
	result := getWorkingDay(Monday)
	fmt.Println(result)
}
