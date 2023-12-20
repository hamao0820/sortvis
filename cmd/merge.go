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
		gui.Run(num, duration, gui.Merge, interactive)
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}
