/*
LINK TO WRAP CALENDAR: https://islipny.gov/community-and-services/wrap-calendar/file
*/
package logic

import "time"

func GetRecyclingDays() map[string][]int {
	var firstDay = 7
	var recyclingType = "paper" // set to first of the year

	days := map[string][]int{
		"glass": {},
		"paper": {},
	}

	// Map through days of the year (up until 365), starting at firstDay
	for i := firstDay; i <= 365; i += 7 {
		// Add to recycling type
		days[recyclingType] = append(days[recyclingType], i)

		// Switch recycling type
		if recyclingType == "paper" {
			recyclingType = "glass"
		} else {
			recyclingType = "paper"
		}
	}

	return days
}

func GetNextRecyclingDay(currentDay int, days []int) int {
	for _, day := range days {
		if day >= currentDay {
			return day
		}
	}
	return -1 // No more recycling days this year
}

// Gets the current day of the year (1-365)
func GetDayOfTheYear() int {
	return time.Now().YearDay()
}
