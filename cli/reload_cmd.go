package cli

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(reloadCmd)
}

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "r",
}
