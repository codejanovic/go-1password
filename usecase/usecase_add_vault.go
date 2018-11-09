package usecase

import (
	"errors"
	"fmt"

	"github.com/codejanovic/gordon/io"
	"github.com/codejanovic/gordon/repository"
)

// AddVaultUsecase usecase
type AddVaultUsecase struct {
}

// AddVaultRequest struct
type AddVaultRequest struct {
	VaultPath  string
	VaultAlias string
}

// AddVaultResponse struct
type AddVaultResponse struct {
	Vault *VaultThinModel
}

var addVaultSingleton *AddVaultUsecase

func init() {
	addVaultSingleton = &AddVaultUsecase{}
}

// NewAddVaultUsecase constructor
func NewAddVaultUsecase() *AddVaultUsecase {
	return addVaultSingleton
}

// Execute the usecase
func (u *AddVaultUsecase) Execute(request *AddVaultRequest) (*AddVaultResponse, error) {
	err := validateAddVaultRequest(request)
	if err != nil {
		return nil, err
	}
	settingsRepository := repository.NewSettingsRepository()
	settings := settingsRepository.Fetch()

	_, err = settings.Find(request.VaultAlias)
	if err == nil {
		return nil, fmt.Errorf("oops, there is already a vault registered with alias '%s'", request.VaultAlias)
	}

	vaultFile := io.NewFileByAbsolutePath(request.VaultPath)
	if !vaultFile.Exists() {
		return nil, fmt.Errorf("oops, it looks like the vault does not exists under path '%s'", request.VaultPath)
	}

	vaultSetting, err := settings.Add(vaultFile.Path(), request.VaultAlias)
	if err != nil {
		return nil, err
	}

	settingsRepository.Store(settings)

	return &AddVaultResponse{
		Vault: &VaultThinModel{
			Alias:      vaultSetting.Alias(),
			Identifier: vaultSetting.Identifier(),
		},
	}, nil
}

func validateAddVaultRequest(request *AddVaultRequest) error {
	if request.VaultPath == "" {
		return errors.New("Please provide a valid path to the opvault file")
	}

	if request.VaultAlias == "" {
		return errors.New("Please provide a unique vault alias to quickly navigate through your vaults")
	}

	return nil
}
