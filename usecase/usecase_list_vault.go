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

// Response struct
type Response struct {
	Found []*ListVault
}

//HasVaults returns whether or not vaults have been found
func (r *Response) HasVaults() bool {
	return len(r.Found) > 0
}

var singleton *ListVaultUsecase

func init() {
	singleton = &ListVaultUsecase{}
}

// NewListVaultUsecase constructor
func NewListVaultUsecase() *ListVaultUsecase {
	return singleton
}

// Execute the usecase
func (u *ListVaultUsecase) Execute() *Response {
	settingsRepository := repository.NewSettingsRepositoryYaml()
	settings := settingsRepository.Fetch()
	vaults := settings.Vaults()
	var found []*ListVault
	for _, vault := range vaults {
		found = append(found, &ListVault{
			Alias: vault.Alias(),
		})
	}
	return &Response{
		Found: found,
	}
}
