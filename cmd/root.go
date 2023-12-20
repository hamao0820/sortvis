package cmd

import (
	"os"

	"github.com/hamao0820/sortvis/gui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sortvis",
	Short: "Visualize sorting algorithms",
	Long:  `Visualize sorting algorithms`,
	RunE: func(cmd *cobra.Command, args []string) error {
		gui.Run(false)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
