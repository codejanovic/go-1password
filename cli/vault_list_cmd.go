package cli

import (
	"encoding/json"
	"github.com/codejanovic/gordon/usecase"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(listVaultCmd)
}

var listVaultCmd = &cobra.Command{
	Use:   "list",
	Short: "list all configured vaults",
	Run: func(cmd *cobra.Command, args []string) {
		response := usecase.NewListVaultUsecase().Execute()
		if !response.HasVaults() {
			cmd.Println("no configured vaults found")
			return
		}

		data, err := json.Marshal(response.Vaults)
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		cmd.Println(string(data))
	},
}
