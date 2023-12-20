package cmd

import (
	"github.com/hamao0820/sortvis/gui"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Visualize merge sort",
	Long:  `Visualize merge sort`,
	Run: func(cmd *cobra.Command, args []string) {
		if num <= 0 || num >= 100 {
			cobra.CheckErr("num must be between 1 and 99")
		}
		if duration <= 0 {
			cobra.CheckErr("duration must be greater than 0")
		}

		gui.Run(num, duration, gui.Merge, interactive)
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}
