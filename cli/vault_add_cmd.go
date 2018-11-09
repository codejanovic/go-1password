package cli

import (
	"encoding/json"

	"github.com/codejanovic/gordon/usecase"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(addVaultCmd)
}

var addVaultCmd = &cobra.Command{
	Use:   "add [alias] [path]",
	Short: "configure a new vault",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		response, err := usecase.NewAddVaultUsecase().Execute(&usecase.AddVaultRequest{
			VaultPath:  args[1],
			VaultAlias: args[0],
		})
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		data, err := json.Marshal(response.Vault)
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		cmd.Println("successfully added vault " + string(data))
	},
}
