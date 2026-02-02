package cmd

import (
	"ko-commands/logic"

	"github.com/spf13/cobra"
)

// recyclingCmd represents the recycling command
var recyclingCmd = &cobra.Command{
	Use:   "recycling",
	Short: "Shows the upcoming recycling",
	Long: `Recycling-related commands:

Run without arguments to get the upcoming recycling day.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDay := logic.GetDayOfTheYear()
		days := logic.GetRecyclingDays()

		nextPaperDay := logic.GetNextRecyclingDay(currentDay, days["paper"])
		nextGlassDay := logic.GetNextRecyclingDay(currentDay, days["glass"])

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
}
