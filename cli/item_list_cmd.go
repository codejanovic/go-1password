package cli

import (
	"encoding/json"
	"github.com/codejanovic/go-1password/usecase"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(listItemCmd)
}

var listItemCmd = &cobra.Command{
	Use:   "list",
	Short: "list all items within a vault",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := usecase.NewListItemsUsecase().Execute()
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		data, err := json.Marshal(response.Items)
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		cmd.Println(string(data))
	},
}
