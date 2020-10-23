package cli

import (
	e "github.com/maximiliangeorg99/SetupHero/environment"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     e.Environment.ProjectName,
	Version: e.Environment.ProjectVersion,
	Short:   e.Environment.ProjectName + " is a simple CLI to make it easy to Setup your local development environment",
	Long:    "Documentation is available at " + e.Environment.ProjectUrl,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
