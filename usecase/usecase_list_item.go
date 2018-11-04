package usecase

import (
	"fmt"
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
	profile, err := requiresActiveProfile()
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
