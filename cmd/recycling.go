/*
LINK TO WRAP CALENDAR: https://islipny.gov/community-and-services/wrap-calendar/file
*/
package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

// recyclingCmd represents the recycling command
var recyclingCmd = &cobra.Command{
	Use:   "recycling",
	Short: "Shows the upcoming recycling",
	Long: `Recycling-related commands:

Run without arguments to get the upcoming recycling day.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDay := getDayOfTheYear()
		days := getRecyclingDays()

		nextPaperDay := getNextRecyclingDay(currentDay, days["paper"])
		nextGlassDay := getNextRecyclingDay(currentDay, days["glass"])

		// Find the closer one
		if nextPaperDay != -1 && (nextGlassDay == -1 || nextPaperDay < nextGlassDay) {
			numberOfDaysUntil := nextPaperDay - currentDay
			cmd.Printf("PAPER in %d days\n", numberOfDaysUntil)
			return
		}
		if nextGlassDay != -1 && (nextPaperDay == -1 || nextGlassDay < nextPaperDay) {
			numberOfDaysUntil := nextGlassDay - currentDay
			cmd.Printf("GLASS/METAL/PLASTIC in %d days\n", numberOfDaysUntil)
			return
		}

		cmd.Println("No more recycling days this year.")
	},
}

func init() {
	rootCmd.AddCommand(recyclingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recyclingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recyclingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getRecyclingDays() map[string][]int {
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

func getNextRecyclingDay(currentDay int, days []int) int {
	for _, day := range days {
		if day >= currentDay {
			return day
		}
	}
	return -1 // No more recycling days this year
}

// Gets the current day of the year (1-365)
func getDayOfTheYear() int {
	return time.Now().YearDay()
}
