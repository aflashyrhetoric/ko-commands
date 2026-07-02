package cmd

import (
	"ko-commands/logic"
	"os"

	"github.com/spf13/cobra"
)

var vanguardCmd = &cobra.Command{
	Use:   "fetch-vanguard-login",
	Short: "Fetch Vanguard login page screenshot",
	Run: func(cmd *cobra.Command, args []string) {
		screenshot, err := logic.FetchVanguardLogin()
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			return
		}
		err = os.WriteFile("vanguard-login.png", screenshot, 0644)
		if err != nil {
			cmd.PrintErrf("Error saving file: %v\n", err)
			return
		}
		cmd.Println("Screenshot saved to vanguard-login.png")
	},
}

func init() {
	rootCmd.AddCommand(vanguardCmd)
}
