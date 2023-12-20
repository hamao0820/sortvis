package cmd

import (
	"os"

	"github.com/hamao0820/sortvis/gui"
	"github.com/spf13/cobra"
)

var (
	interactive bool
	num         int
)

var rootCmd = &cobra.Command{
	Use:   "sortvis",
	Short: "Visualize sorting algorithms",
	Long:  `Visualize sorting algorithms`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if num <= 0 || num >= 100 {
			cobra.CheckErr("num must be between 0 and 99")
		}
		gui.Run(num, interactive)
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
	rootCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "interactive mode")
	rootCmd.Flags().IntVarP(&num, "num", "n", 50, "number of elements")
}
