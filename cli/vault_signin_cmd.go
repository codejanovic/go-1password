package cli

import (
	"github.com/codejanovic/go-1password/usecase"
	"github.com/spf13/cobra"
)

func init() {
	vaultCmd.AddCommand(signinVaultCmd)
	addVaultCmd.Flags().StringP("secret", "s", "", "provide a vault secret to open the vault")
	addVaultCmd.Flags().StringP("profile", "p", "", "provide a profile to signin")
}

var signinVaultCmd = &cobra.Command{
	Use:   "signin",
	Short: "sign into a vault",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := usecase.NewSignInVaultUsecase().Execute(&usecase.SignInVaultRequest{
			VaultAliasOrIdentifier: args[0],
			VaultProfile:           cmd.Flag("profile").Value.String(),
			VaultSecret:            cmd.Flag("secret").Value.String(),
		})
		if err != nil {
			cmd.Println(err.Error())
			return
		}
		cmd.Printf("successfully signed into vault '%s' with profile '%s'", args[0], args[1])
	},
}
