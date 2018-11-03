package usecase

import repository "github.com/codejanovic/go-1password/repository"

// RemoveVaultUsecase usecase
type RemoveVaultUsecase struct {
}

// RemoveVaultRequest struct
type RemoveVaultRequest struct {
	VaultAliasOrIdentifier string
}

var removeVaultSingleton *RemoveVaultUsecase

func init() {
	removeVaultSingleton = &RemoveVaultUsecase{}
}

// NewRemoveVaultUsecase constructor
func NewRemoveVaultUsecase() *RemoveVaultUsecase {
	return removeVaultSingleton
}

// Execute usecase
func (u *RemoveVaultUsecase) Execute(request *RemoveVaultRequest) error {
	settingsRepository := repository.NewSettingsRepository()
	credentialsRepository := repository.NewCredentialsRepository()
	settings := settingsRepository.Fetch()
	vault, err := settings.Find(request.VaultAliasOrIdentifier)
	if err != nil {
		return err
	}

	settings.Remove(vault.Identifier())
	credentialsRepository.Remove(vault.Identifier())
	settingsRepository.Store(settings)
	return nil
}
