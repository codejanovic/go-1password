package usecase

import (
	"errors"
	"fmt"
	"strings"
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

	profile, err := requiresActiveProfile()
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
