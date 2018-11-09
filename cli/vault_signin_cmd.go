package cli

import (
	"github.com/codejanovic/gordon/usecase"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(signinVaultCmd)
	addVaultCmd.Flags().StringP("secret", "s", "", "secret to open the vault/profile")
}

var signinVaultCmd = &cobra.Command{
	Use:   "signin [alias|identifier] [profile]",
	Short: "sign into a vault profile by default",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := usecase.NewSignInVaultUsecase().Execute(&usecase.SignInVaultRequest{
			VaultAliasOrIdentifier: args[0],
			VaultProfile:           args[1],
			VaultSecret:            cmd.Flag("secret").Value.String(),
		})
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		cmd.Printf("successfully signed into vault '%s' with profile '%s'", args[0], args[1])
	},
}
