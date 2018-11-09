package usecase

import (
	"fmt"

	"github.com/codejanovic/gordon/model"

	"github.com/codejanovic/gordon/interaction"
	"github.com/codejanovic/gordon/repository"
	"github.com/codejanovic/gordon/vault"
)

func requireActiveOrAlternativeVault(alternativeVault string) (vault.Vault, model.VaultSetting, error) {
	settingsRepository := repository.NewSettingsRepository()
	settings := settingsRepository.Fetch()

	vaultSetting, err := settings.ActiveOrAlternative(alternativeVault)
	if err != nil {
		return nil, nil, err
	}

	foundVault := vault.NewOpVault(vaultSetting.Path(), vaultSetting.Profile())
	return foundVault, vaultSetting, nil
}

func requireActiveOrAlternativeProfile(alternativeVault string, alternativeProfile string) (vault.Profile, error) {
	credentialsRepository := repository.NewCredentialsRepository()

	foundVault, foundVaultSetting, err := requireActiveOrAlternativeVault(alternativeVault)
	if err != nil {
		return nil, fmt.Errorf("unable to find active vault. did you sign in into a default vault and profile? ")
	}

	if alternativeProfile == "" {
		if !foundVault.HasDefaultProfile() {
			return nil, fmt.Errorf("unable to open vault profile. did you sign in into a default vault and profile? ")
		}
		foundProfile, err := requiresOpenedProfile(
			func(secret string) (vault.Profile, error) {
				return foundVault.OpenDefaultProfile(secret)
			},
			func() (string, bool) {
				return credentialsRepository.Fetch(foundVaultSetting.Identifier())
			},
			func(secret string) {
				credentialsRepository.Store(foundVaultSetting.Identifier(), secret)
			})

		return foundProfile, err

	}

	foundProfile, err := requiresOpenedProfile(
		func(secret string) (vault.Profile, error) {
			return foundVault.OpenProfile(alternativeProfile, secret)
		},
		func() (string, bool) {
			return credentialsRepository.Fetch(foundVaultSetting.Identifier())
		},
		func(secret string) {
			credentialsRepository.Store(foundVaultSetting.Identifier(), secret)
		})

	return foundProfile, err
}

func requiresOpenedProfile(login func(secret string) (vault.Profile, error), fetchSecret func() (string, bool), storeSecret func(secret string)) (vault.Profile, error) {
	secretFromCredentialStore, found := fetchSecret()
	var profile vault.Profile
	var err error
	if found {
		profile, err = login(secretFromCredentialStore)
		if err == nil {
			return profile, nil
		}
	}

	userAction := interaction.NewConsoleInteraction()
	for i := 0; i < 3; i++ {
		secretFromUser, err := userAction.AskForVaultPassword()
		if err != nil {
			continue
		}
		profile, err = login(secretFromUser)
		if err != nil {
			continue
		}

		storeSecret(secretFromUser)
		return profile, nil
	}

	return nil, fmt.Errorf("unable to open vault profile. did you provide the correct secret? ")
}
