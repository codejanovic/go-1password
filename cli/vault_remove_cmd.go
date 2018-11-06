package cli

import (
	"github.com/codejanovic/go-1password/usecase"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(removeVaultCmd)
	removeVaultCmd.Flags().StringP("file", "f", "", "provide a valid path to your opvault")
	removeVaultCmd.Flags().StringP("alias", "a", "", "provide a unique vault alias")
	removeVaultCmd.MarkFlagRequired("file")
	removeVaultCmd.MarkFlagRequired("alias")
}

var removeVaultCmd = &cobra.Command{
	Use:   "remove",
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
