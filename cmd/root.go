package cmd

import (
	"os"

	g "github.com/hamao0820/sortvis/graph"
	"github.com/hamao0820/sortvis/ui"
	"github.com/hamao0820/sortvis/util"
	"github.com/spf13/cobra"
)

var (
	interactive bool
	num         int
	duration    int
	file        string
	graph       string
	ls          bool
	graphList   bool
)

var rootCmd = &cobra.Command{
	Use:     "sortvis",
	Version: "v0.4.1",
	Short:   "Visualize sorting algorithms",
	Long: `Visualize sorting algorithms.
in interactive mode, you can step forward by pressing space.
Press q or Ctrl+C to quit`,
	Run: func(cmd *cobra.Command, args []string) {
		if ls {
			cmd.Println("You can use the following algorithms(subcommands):")
			for _, v := range []ui.Algorithm{ui.Bubble, ui.Merge, ui.Heap, ui.Quick, ui.Selection, ui.Bucket, ui.Insertion, ui.Shell} {
				cmd.Println(v.Pretty() + "(" + v.String() + ")")
			}
			return
		}

		if graphList {
			cmd.Println("You can use the following graph types:")
			for _, g := range g.GraphList {
				cmd.Println(g.Name + ": " + g.Description)
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
	rootCmd.PersistentFlags().StringVarP(&graph, "graph", "g", "", `The initial value of the array is set to a certain form. Example: sin wave
	You can check the available list with 'sortvis --graph-list'.`)
	rootCmd.Flags().BoolVarP(&ls, "ls", "l", false, "list all algorithms")
	rootCmd.Flags().BoolVar(&graphList, "graph-list", false, "list all graph types")
}
