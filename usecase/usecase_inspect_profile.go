package usecase

import (
	"fmt"
)

// InspectProfileUsecase usecase
type InspectProfileUsecase struct {
}

// InspectProfileRequest struct
type InspectProfileRequest struct {
	AlternativeVault   string
	AlternativeProfile string
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
func (u *InspectProfileUsecase) Execute(request *InspectProfileRequest) (*InspectProfileResponse, error) {
	profile, err := requireActiveOrAlternativeProfile(request.AlternativeVault, request.AlternativeProfile)
	if err != nil {
		return nil, fmt.Errorf("unable to open vault profile. did you sign in into a default vault and profile? ")
	}

	return &InspectProfileResponse{
		Profile: toProfileInspectModel(profile),
	}, nil
}
