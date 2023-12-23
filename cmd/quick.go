package cmd

import (
	"github.com/hamao0820/sortvis/ui"
	"github.com/spf13/cobra"
)

var quickCmd = &cobra.Command{
	Use:     "quick",
	Version: rootCmd.Version,
	Short:   "Visualize quick sort",
	Long:    `Visualize quick sort`,
	Run: func(cmd *cobra.Command, args []string) {
		if num <= 0 || num >= 100 {
			cobra.CheckErr("num must be between 1 and 99")
		}
		if duration <= 0 {
			cobra.CheckErr("duration must be greater than 0")
		}

		err := ui.Run(num, duration, ui.Quick, file, graph, interactive)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(quickCmd)
}
