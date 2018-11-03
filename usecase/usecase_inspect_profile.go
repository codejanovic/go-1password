package usecase

import (
	"errors"
	"fmt"

	repository "github.com/codejanovic/go-1password/repository"
	vault "github.com/codejanovic/go-1password/vault"
)

// InspectProfileUsecase usecase
type InspectProfileUsecase struct {
}

// InspectProfileRequest struct
type InspectProfileRequest struct {
	VaultAliasOrIdentifier string
}

// InspectProfile struct
type InspectProfile struct {
	Name  string         `json:"name"`
	Items []*InspectItem `json:"items"`
}

// InspectItem struct
type InspectItem struct {
	Name   string `json:"name"`
	Fields int    `json:"fields"`
}

// InspectProfileResponse struct
type InspectProfileResponse struct {
	Profile *InspectProfile
}

var inspectProfileSingleton *InspectProfileUsecase

func init() {
	inspectProfileSingleton = &InspectProfileUsecase{}
}

// NewInspectProfileUsecase constructor
func NewInspectProfileUsecase() *InspectProfileUsecase {
	return inspectProfileSingleton
}

// Execute the usecase
func (u *InspectProfileUsecase) Execute() (*InspectProfileResponse, error) {
	settingsRepository := repository.NewSettingsRepository()
	settings := settingsRepository.Fetch()
	credentialsRepository := repository.NewCredentialsRepository()

	vaultSetting, err := settings.Active()
	if err != nil {
		return nil, fmt.Errorf("Unable to find active vault. Please signin first")
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

	inspectItems := make([]*InspectItem, 0)
	for _, item := range profile.Items() {
		inspectItems = append(inspectItems, &InspectItem{
			Name:   item.Name(),
			Fields: item.FieldSize(),
		})
	}

	return &InspectProfileResponse{
		Profile: &InspectProfile{
			Name:  profile.Name(),
			Items: inspectItems,
		},
	}, nil
}
