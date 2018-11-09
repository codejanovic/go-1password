package cli

import (
	"encoding/json"

	"github.com/codejanovic/gordon/usecase"
	"github.com/spf13/cobra"
)

func init() {
	itemCmd.AddCommand(listItemCmd)
	listItemCmd.Flags().StringP("vault", "v", "", "alternative vault")
	listItemCmd.Flags().StringP("profile", "p", "", "alternative profile")
}

var listItemCmd = &cobra.Command{
	Use:   "list",
	Short: "list all items within a vault",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		response, err := usecase.NewListItemsUsecase().Execute(&usecase.ListItemRequest{
			AlternativeVault:   cmd.Flag("vault").Value.String(),
			AlternativeProfile: cmd.Flag("profile").Value.String(),
		})
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
