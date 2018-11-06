package cli

import (
	"fmt"
	e "github.com/codejanovic/go-1password/environment"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + e.Environment.ProjectName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(e.Environment.ProjectName + " v" + e.Environment.ProjectVersion)
	},
}
