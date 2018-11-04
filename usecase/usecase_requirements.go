package usecase

import (
	"errors"
	"fmt"

	repository "github.com/codejanovic/go-1password/repository"
	vault "github.com/codejanovic/go-1password/vault"
)

func requiresActiveVault() (vault.Vault, error) {
	settingsRepository := repository.NewSettingsRepository()
	settings := settingsRepository.Fetch()

	vaultSetting, err := settings.Active()
	if err != nil {
		return nil, fmt.Errorf("Unable to find active vault. Please signin first")
	}
	vault := vault.NewOpVault(vaultSetting.Path())
	return vault, nil
}

func requiresActiveProfile() (vault.Profile, error) {
	settingsRepository := repository.NewSettingsRepository()
	settings := settingsRepository.Fetch()
	credentialsRepository := repository.NewCredentialsRepository()

	vaultSetting, err := settings.Active()
	if err != nil {
		return nil, fmt.Errorf("Unable to find active vault. Make sure to sign first")
	}
	secret, found := credentialsRepository.Fetch(vaultSetting.Identifier())
	if !found {
		return nil, errors.New("Unable to find credentials. Make sure to sign first")
	}

	vault := vault.NewOpVault(vaultSetting.Path())
	profile, err := vault.OpenProfile(vaultSetting.Profile(), secret)
	if err != nil {
		return nil, fmt.Errorf("Unable to open vault profile. Make sure to sign first")
	}
	return profile, nil
}
