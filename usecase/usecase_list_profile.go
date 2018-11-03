package usecase

import (
	"fmt"

	repository "github.com/codejanovic/go-1password/repository"
	vault "github.com/codejanovic/go-1password/vault"
)

// ListProfileUsecase usecase
type ListProfileUsecase struct {
}

// ListProfileRequest struct
type ListProfileRequest struct {
	VaultAliasOrIdentifier string
}

// ListProfile struct
type ListProfile struct {
	Alias      string
	Identifier string
}

// ListProfileResponse struct
type ListProfileResponse struct {
	Found []string
}

//HasVaults returns whether or not vaults have been found
func (r *ListProfileResponse) HasVaults() bool {
	return len(r.Found) > 0
}

var listProfileSingleton *ListProfileUsecase

func init() {
	listProfileSingleton = &ListProfileUsecase{}
}

// NewListProfileUsecase constructor
func NewListProfileUsecase() *ListProfileUsecase {
	return listProfileSingleton
}

// Execute the usecase
func (u *ListProfileUsecase) Execute() (*ListProfileResponse, error) {
	settingsRepository := repository.NewSettingsRepository()
	settings := settingsRepository.Fetch()

	vaultSetting, err := settings.Active()
	if err != nil {
		return nil, fmt.Errorf("Unable to find active vault. Please signin first")
	}
	vault := vault.NewOpVault(vaultSetting.Path())
	foundProfiles, err := vault.Profiles()
	if err != nil {
		return nil, err
	}

	return &ListProfileResponse{
		Found: foundProfiles,
	}, nil
}
