package usecase

import (
	"github.com/codejanovic/gordon/repository"
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
	vaultSettings := settings.Vaults()

	return &ListVaultResponse{
		Vaults: toVaultThinModels(vaultSettings),
	}
}
