package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(itemCmd)
}

var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "item actions",
}
