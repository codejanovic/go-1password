package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(profileCmd)
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "profile actions",
}
