package cmd

import (
	"github.com/hamao0820/sortvis/ui"
	"github.com/spf13/cobra"
)

var heapCmd = &cobra.Command{
	Use:     "heap",
	Version: rootCmd.Version,
	Short:   "Visualize heap sort",
	Long:    `Visualize heap sort`,
	Run: func(cmd *cobra.Command, args []string) {
		if num <= 0 || num >= 100 {
			cobra.CheckErr("num must be between 1 and 99")
		}
		if duration <= 0 {
			cobra.CheckErr("duration must be greater than 0")
		}

		err := ui.Run(num, duration, ui.Heap, file, graph, interactive)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(heapCmd)
}
