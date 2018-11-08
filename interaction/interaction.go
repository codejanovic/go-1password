package interaction

type Interaction interface {
	AskForPassword(command string) (string, error)
	AskForVaultPassword() (string, error)
	AskForConfirmation(question string) (string, error)
}
