package cli

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(cleanupdumpsCmd)
}

var cleanupdumpsCmd = &cobra.Command{
	Use:   "clean",
	Short: "c",
}
