package cmd

import (
	"github.com/hamao0820/sortvis/gui"
	"github.com/spf13/cobra"
)

var heapCmd = &cobra.Command{
	Use:   "heap",
	Short: "Visualize heap sort",
	Long:  `Visualize heap sort`,
	Run: func(cmd *cobra.Command, args []string) {
		gui.Run(num, duration, gui.Heap, interactive)
	},
}

func init() {
	rootCmd.AddCommand(heapCmd)
}
