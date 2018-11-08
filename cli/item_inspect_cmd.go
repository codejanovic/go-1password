package cli

import (
	"encoding/json"
	"github.com/codejanovic/gordon/usecase"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(inspectItemCmd)
}

var inspectItemCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect item",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := usecase.NewInspectItemUsecase().Execute(&usecase.InspectItemRequest{
			ItemName: args[0],
		})
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		data, err := json.Marshal(response.Item)
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		cmd.Println(string(data))
	},
}
