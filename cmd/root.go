package cmd

import (
	"os"

	"github.com/hamao0820/sortvis/gui"
	"github.com/hamao0820/sortvis/util"
	"github.com/spf13/cobra"
)

var (
	interactive bool
	num         int
	duration    int
	file        string
	ls          bool
)

var rootCmd = &cobra.Command{
	Use:   "sortvis",
	Short: "Visualize sorting algorithms",
	Long: `Visualize sorting algorithms.
algorithms: bubble, merge, heap
in interactive mode, you can step forward
press q or Ctrl+C to quit`,
	Run: func(cmd *cobra.Command, args []string) {
		if ls {
			cmd.Println("You can use the following algorithms(subcommands):")
			for _, v := range []gui.Algorithm{gui.Bubble, gui.Merge, gui.Heap} {
				cmd.Println(v.Pretty() + "(" + v.String() + ")")
			}
			return
		}
		cmd.Println(util.Logo)
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
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "file path: number of lines must be 0 < n < 100, each line must be 0 <= n <= 100")
	rootCmd.Flags().BoolVarP(&ls, "ls", "l", false, "list all algorithms")
}
