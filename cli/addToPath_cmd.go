package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addToPathCmd)
}

var addToPathCmd = &cobra.Command{
	Use:   "addToPath",
	Short: "a",
	Args:  cobra.MinimumNArgs(2),
	Run:   func(cmd *cobra.Command, args []string) {},
}
