package cli

import (
	"encoding/json"

	"github.com/codejanovic/gordon/usecase"
	"github.com/spf13/cobra"
)

func init() {
	itemCmd.AddCommand(inspectItemCmd)
	inspectItemCmd.Flags().StringP("vault", "v", "", "alternative vault")
	inspectItemCmd.Flags().StringP("profile", "p", "", "alternative profile")
}

var inspectItemCmd = &cobra.Command{
	Use:   "inspect [name]",
	Short: "inspect item",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response, err := usecase.NewInspectItemUsecase().Execute(&usecase.InspectItemRequest{
			ItemName:           args[0],
			AlternativeVault:   cmd.Flag("vault").Value.String(),
			AlternativeProfile: cmd.Flag("profile").Value.String(),
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
