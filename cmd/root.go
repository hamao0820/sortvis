package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	interactive bool
	num         int
	duration    int
)

var rootCmd = &cobra.Command{
	Use:   "sortvis",
	Short: "Visualize sorting algorithms",
	Long:  `Visualize sorting algorithms`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if num <= 0 || num >= 100 {
			cobra.CheckErr("num must be between 0 and 99")
		}
		if duration <= 0 {
			cobra.CheckErr("duration must be greater than 0")
		}

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
	rootCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "interactive mode")
	rootCmd.PersistentFlags().IntVarP(&num, "num", "n", 50, "number of elements")
	rootCmd.PersistentFlags().IntVarP(&duration, "duration", "d", 300, "duration of each step in milliseconds")
}
