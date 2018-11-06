package cli

import (
	"encoding/json"
	"github.com/codejanovic/go-1password/usecase"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(inspectProfileCmd)
}

var inspectProfileCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspect profile",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := usecase.NewInspectProfileUsecase().Execute()
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
