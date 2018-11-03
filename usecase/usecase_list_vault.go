package usecase

import (
	repository "github.com/codejanovic/go-1password/repository"
)

// ListVaultUsecase usecase
type ListVaultUsecase struct {
}

// ListVault struct
type ListVault struct {
	Alias string
}

// ListVaultResponse struct
type ListVaultResponse struct {
	Found []*ListVault
}

//HasVaults returns whether or not vaults have been found
func (r *ListVaultResponse) HasVaults() bool {
	return len(r.Found) > 0
}

var listVaultSingleton *ListVaultUsecase

func init() {
	listVaultSingleton = &ListVaultUsecase{}
}

// NewListVaultUsecase constructor
func NewListVaultUsecase() *ListVaultUsecase {
	return listVaultSingleton
}

// Execute the usecase
func (u *ListVaultUsecase) Execute() *ListVaultResponse {
	settingsRepository := repository.NewSettingsRepositoryYaml()
	settings := settingsRepository.Fetch()
	vaults := settings.Vaults()
	var found []*ListVault
	for _, vault := range vaults {
		found = append(found, &ListVault{
			Alias: vault.Alias(),
		})
	}
	return &ListVaultResponse{
		Found: found,
	}
}
