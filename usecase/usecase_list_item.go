package usecase

import (
	"errors"
	"fmt"

	repository "github.com/codejanovic/go-1password/repository"
	vault "github.com/codejanovic/go-1password/vault"
)

// ListItemsUsecase usecase
type ListItemsUsecase struct {
}

// ListItemsResponse struct
type ListItemsResponse struct {
	Items []*ItemThinModel
}

var listItemsSingleton *ListItemsUsecase

func init() {
	listItemsSingleton = &ListItemsUsecase{}
}

// NewListItemsUsecase constructor
func NewListItemsUsecase() *ListItemsUsecase {
	return listItemsSingleton
}

// Execute the usecase
func (u *ListItemsUsecase) Execute() (*ListItemsResponse, error) {
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

	listItems := make([]*ItemThinModel, 0)
	for _, item := range profile.Items() {
		listItems = append(listItems, &ItemThinModel{
			Name:   item.Name(),
			Fields: item.FieldSize(),
		})
	}

	return &ListItemsResponse{
		Items: listItems,
	}, nil
}
