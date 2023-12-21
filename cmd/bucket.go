package cmd

import (
	"github.com/hamao0820/sortvis/gui"
	"github.com/spf13/cobra"
)

var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Visualize bucket sort",
	Long:  `Visualize bucket sort`,
	Run: func(cmd *cobra.Command, args []string) {
		if num <= 0 || num >= 100 {
			cobra.CheckErr("num must be between 1 and 99")
		}
		if duration <= 0 {
			cobra.CheckErr("duration must be greater than 0")
		}

		err := gui.Run(num, duration, gui.Bucket, file, interactive)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(bucketCmd)
}