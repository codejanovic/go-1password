package cli

import (
	"encoding/json"

	"github.com/codejanovic/gordon/usecase"
	"github.com/spf13/cobra"
)

func init() {
	profileCmd.AddCommand(inspectProfileCmd)
	inspectProfileCmd.Flags().StringP("vault", "v", "", "alternative vault")
	inspectProfileCmd.Flags().StringP("profile", "p", "", "alternative profile")
}

var inspectProfileCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect profile",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		response, err := usecase.NewInspectProfileUsecase().Execute(
			&usecase.InspectProfileRequest{
				AlternativeVault:   cmd.Flag("vault").Value.String(),
				AlternativeProfile: cmd.Flag("profile").Value.String(),
			})
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		data, err := json.Marshal(response.Profile)
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		cmd.Println(string(data))
	},
}
