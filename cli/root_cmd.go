package cli

import (
	"fmt"
	e "github.com/codejanovic/go-1password/environment"
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
		fmt.Println(err)
		os.Exit(1)
	}
}
