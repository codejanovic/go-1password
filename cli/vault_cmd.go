package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(vaultCmd)
}

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "vault actions",
}
