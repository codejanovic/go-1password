package usecase

import (
	repository "github.com/codejanovic/go-1password/repository"
)

// ListVaultUsecase usecase
type ListVaultUsecase struct {
}

// ListVaultResponse struct
type ListVaultResponse struct {
	Vaults []*VaultThinModel
}

//HasVaults returns whether or not vaults have been found
func (r *ListVaultResponse) HasVaults() bool {
	return len(r.Vaults) > 0
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
	settingsRepository := repository.NewSettingsRepository()
	settings := settingsRepository.Fetch()
	vaults := settings.Vaults()
	var foundVaults []*VaultThinModel
	for _, vault := range vaults {
		foundVaults = append(foundVaults, &VaultThinModel{
			Alias:      vault.Alias(),
			Identifier: vault.Identifier(),
		})
	}
	return &ListVaultResponse{
		Vaults: foundVaults,
	}
}
