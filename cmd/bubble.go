package cmd

import (
	"github.com/hamao0820/sortvis/gui"
	"github.com/spf13/cobra"
)

var bubbleCmd = &cobra.Command{
	Use:   "bubble",
	Short: "Visualize bubble sort",
	Long:  `Visualize bubble sort`,
	Run: func(cmd *cobra.Command, args []string) {
		gui.Run(num, duration, gui.Bubble, interactive)
	},
}

func init() {
	rootCmd.AddCommand(bubbleCmd)
}
