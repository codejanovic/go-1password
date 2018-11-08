package usecase

import (
	"errors"
	"fmt"

	"github.com/codejanovic/gordon/repository"
	"github.com/codejanovic/gordon/vault"
)

// SignInVaultUsecase usecase
type SignInVaultUsecase struct {
}

// SignInVaultRequest struct
type SignInVaultRequest struct {
	VaultAliasOrIdentifier string
	VaultProfile           string
	VaultSecret            string
}

var signInVaultSingleton *SignInVaultUsecase

func init() {
	signInVaultSingleton = &SignInVaultUsecase{}
}

// NewSignInVaultUsecase constructor
func NewSignInVaultUsecase() *SignInVaultUsecase {
	return signInVaultSingleton
}

// Execute usecase
func (u *SignInVaultUsecase) Execute(request *SignInVaultRequest) error {
	err := validateSignInVaultRequest(request)
	if err != nil {
		return err
	}

	settingsRepository := repository.NewSettingsRepository()
	credentialsRepository := repository.NewCredentialsRepository()
	settings := settingsRepository.Fetch()
	vaultSetting, err := settings.Find(request.VaultAliasOrIdentifier)
	if err != nil {
		return err
	}

	var secret string
	if request.VaultSecret == "" {
		s, found := credentialsRepository.Fetch(vaultSetting.Identifier())
		if !found {
			return fmt.Errorf("Unable to find secret within the credentials store. Please provide a vault secret manually")
		}
		secret = s
	} else {
		secret = request.VaultSecret
	}

	settings.Activate(vaultSetting.Identifier())
	settingsRepository.Store(settings)

	vault := vault.NewOpVault(vaultSetting.Path())
	_, err = vault.OpenProfile(request.VaultProfile, secret)
	if err != nil {
		return err
	}

	vaultSetting.WithProfile(request.VaultProfile)
	settingsRepository.Store(settings)
	credentialsRepository.Store(vaultSetting.Identifier(), secret)
	return nil
}

func validateSignInVaultRequest(request *SignInVaultRequest) error {
	if request.VaultAliasOrIdentifier == "" {
		return errors.New("Please provide a vault alias or identifier")
	}

	if request.VaultProfile == "" {
		return errors.New("Please provide a vault profile")
	}

	return nil
}
