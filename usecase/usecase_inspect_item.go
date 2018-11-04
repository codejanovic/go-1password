package usecase

import (
	"errors"
	"fmt"
	"strings"

	repository "github.com/codejanovic/go-1password/repository"
	vault "github.com/codejanovic/go-1password/vault"
)

// InspectItemUsecase usecase
type InspectItemUsecase struct {
}

// InspectItemRequest struct
type InspectItemRequest struct {
	ItemName string
}

// InspectItemResponse struct
type InspectItemResponse struct {
	Item *ItemInspectModel
}

var inspectItemSingleton *InspectItemUsecase

func init() {
	inspectItemSingleton = &InspectItemUsecase{}
}

// NewInspectItemUsecase constructor
func NewInspectItemUsecase() *InspectItemUsecase {
	return inspectItemSingleton
}

// Execute the usecase
func (u *InspectItemUsecase) Execute(request *InspectItemRequest) (*InspectItemResponse, error) {
	if request.ItemName == "" {
		return nil, errors.New("Please provide a item name you want to inspect")
	}

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

	var inspectItem *ItemInspectModel
	inspectFields := make([]*FieldInspectModel, 0)
	for _, item := range profile.Items() {
		if strings.EqualFold(item.Name(), request.ItemName) {
			for _, field := range item.Fields() {
				inspectFields = append(inspectFields, &FieldInspectModel{
					Name:       field.Name(),
					Value:      field.Value(),
					IsPassword: field.IsPassword(),
				})
			}
			inspectItem = &ItemInspectModel{
				Name:   item.Name(),
				Fields: inspectFields,
			}
		}
	}

	if inspectItem == nil {
		return nil, fmt.Errorf("Unable to find item with name %s", request.ItemName)
	}

	return &InspectItemResponse{
		Item: inspectItem,
	}, nil
}
