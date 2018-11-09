package usecase

import (
	"errors"
	"github.com/codejanovic/gordon/repository"
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
	settings := settingsRepository.Fetch()
	vaultSetting, err := settings.Find(request.VaultAliasOrIdentifier)
	if err != nil {
		return err
	}

	vaultSetting.WithProfile(request.VaultProfile)
	settings.Activate(vaultSetting.Identifier())
	settingsRepository.Store(settings)

	_, err = requireActiveOrAlternativeProfile("", "")
	if err != nil {
		return err
	}

	return nil
}

func validateSignInVaultRequest(request *SignInVaultRequest) error {
	if request.VaultAliasOrIdentifier == "" {
		return errors.New("please provide a vault alias or identifier")
	}

	if request.VaultProfile == "" {
		return errors.New("please provide a vault profile")
	}

	return nil
}
