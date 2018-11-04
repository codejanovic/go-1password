package usecase

// ItemInspectModel struct
type ItemInspectModel struct {
	Name   string               `json:"name"`
	Fields []*FieldInspectModel `json:"fields"`
}

// FieldInspectModel struct
type FieldInspectModel struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	IsPassword bool   `json:"isPassword"`
}

// ProfileInspectModel struct
type ProfileInspectModel struct {
	Name  string           `json:"name"`
	Items []*ItemThinModel `json:"items"`
}

// ItemThinModel struct
type ItemThinModel struct {
	Name   string `json:"name"`
	Fields int    `json:"fields"`
}

// ProfileThinModel struct
type ProfileThinModel struct {
	Name  string `json:"name"`
	Items int    `json:"items"`
}

// ProfileNameOnlyModel struct
type ProfileNameOnlyModel struct {
	Name string `json:"name"`
}

// VaultThinModel struct
type VaultThinModel struct {
	Identifier string `json:"identifier"`
	Alias      string `json:"alias"`
}
