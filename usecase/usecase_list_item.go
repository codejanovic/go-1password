package usecase

import (
	"fmt"
)

// ListItemsUsecase usecase
type ListItemsUsecase struct {
}

// ListItemRequest struct
type ListItemRequest struct {
	AlternativeVault   string
	AlternativeProfile string
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
func (u *ListItemsUsecase) Execute(request *ListItemRequest) (*ListItemsResponse, error) {
	profile, err := requireActiveOrAlternativeProfile(request.AlternativeVault, request.AlternativeProfile)
	if err != nil {
		return nil, fmt.Errorf("Unable to open vault profile. Make sure to sign first")
	}

	return &ListItemsResponse{
		Items: toItemThinModels(profile),
	}, nil
}
