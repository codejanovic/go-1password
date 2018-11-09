package usecase

// ListProfileUsecase usecase
type ListProfileUsecase struct {
}

// ListProfileRequest struct
type ListProfileRequest struct {
	AlternativeVault string
}

// ListProfileResponse struct
type ListProfileResponse struct {
	Profiles []*ProfileNameOnlyModel
}

var listProfileSingleton *ListProfileUsecase

func init() {
	listProfileSingleton = &ListProfileUsecase{}
}

// NewListProfileUsecase constructor
func NewListProfileUsecase() *ListProfileUsecase {
	return listProfileSingleton
}

// Execute the usecase
func (u *ListProfileUsecase) Execute(request *ListProfileRequest) (*ListProfileResponse, error) {
	vault, _, err := requireActiveOrAlternativeVault(request.AlternativeVault)
	if err != nil {
		return nil, err
	}

	foundProfiles, err := vault.Profiles()
	if err != nil {
		return nil, err
	}

	return &ListProfileResponse{
		Profiles: toProfileNameOnlyModels(foundProfiles),
	}, nil
}
