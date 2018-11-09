package cli

import (
	"github.com/codejanovic/gordon/usecase"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(removeVaultCmd)
}

var removeVaultCmd = &cobra.Command{
	Use:   "remove [alias|identifier]",
	Short: "remove vault from configuration",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := usecase.NewRemoveVaultUsecase().Execute(&usecase.RemoveVaultRequest{
			VaultAliasOrIdentifier: args[0],
		})
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		cmd.Printf("successfully removed vault '%s'", args[0])
	},
}
