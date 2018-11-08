package cli

import (
	"encoding/json"
	"github.com/codejanovic/gordon/fatal"
	"github.com/codejanovic/gordon/usecase"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(addVaultCmd)
	addVaultCmd.Flags().StringP("file", "f", "", "provide a valid path to your opvault")
	addVaultCmd.Flags().StringP("alias", "a", "", "provide a unique vault alias")
	err := addVaultCmd.MarkFlagRequired("file")
	if err != nil {
		fatal.Crash(err, "")
	}
	err = addVaultCmd.MarkFlagRequired("alias")
	if err != nil {
		fatal.Crash(err, "")
	}
}

var addVaultCmd = &cobra.Command{
	Use:   "add",
	Short: "configure a new vault",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		response, err := usecase.NewAddVaultUsecase().Execute(&usecase.AddVaultRequest{
			VaultPath:  cmd.Flag("file").Value.String(),
			VaultAlias: cmd.Flag("alias").Value.String(),
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
