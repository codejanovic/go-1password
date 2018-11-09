package cli

import (
	"encoding/json"

	"github.com/codejanovic/gordon/usecase"
	"github.com/spf13/cobra"
)

func init() {
	profileCmd.AddCommand(listProfileCmd)
	listProfileCmd.Flags().StringP("vault", "v", "", "alternative vault")
}

var listProfileCmd = &cobra.Command{
	Use:   "list",
	Short: "list all profiles within a vault",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := usecase.NewListProfileUsecase().Execute(&usecase.ListProfileRequest{
			AlternativeVault: cmd.Flag("vault").Value.String(),
		})
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		data, err := json.Marshal(response.Profiles)
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		cmd.Println(string(data))
	},
}
