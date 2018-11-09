package cli

import (
	e "github.com/codejanovic/gordon/environment"
	"github.com/codejanovic/gordon/fatal"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     e.Environment.ProjectName,
	Version: e.Environment.ProjectVersion,
	Short:   e.Environment.ProjectName + " is a simple CLI for interacting with vaults in opvault format",
	Long:    "Documentation is available at " + e.Environment.ProjectUrl,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fatal.Crash(err, "Lets see what the stacktrace talks")
		os.Exit(1)
	}
}
