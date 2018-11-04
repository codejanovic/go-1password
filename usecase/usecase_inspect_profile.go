package usecase

import (
	"fmt"
)

// InspectProfileUsecase usecase
type InspectProfileUsecase struct {
}

// InspectProfileRequest struct
type InspectProfileRequest struct {
	VaultAliasOrIdentifier string
}

// InspectProfileResponse struct
type InspectProfileResponse struct {
	Profile *ProfileInspectModel
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
	profile, err := requiresActiveProfile()
	if err != nil {
		return nil, fmt.Errorf("Unable to open vault profile. Make sure to sign first")
	}

	return &InspectProfileResponse{
		Profile: toProfileInspectModel(profile),
	}, nil
}
