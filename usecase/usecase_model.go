package usecase

import (
	"github.com/codejanovic/go-1password/model"
	"github.com/codejanovic/go-1password/vault"
)

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

func toFieldInspectModel(field vault.Field) *FieldInspectModel {
	return &FieldInspectModel{
		Name:       field.Name(),
		Value:      field.Value(),
		IsPassword: field.IsPassword(),
	}
}

func toItemInspectModel(item vault.Item) *ItemInspectModel {
	var inspectItem *ItemInspectModel
	inspectFields := make([]*FieldInspectModel, 0)
	for _, field := range item.Fields() {
		inspectFields = append(inspectFields, toFieldInspectModel(field))
	}
	inspectItem = &ItemInspectModel{
		Name:   item.Name(),
		Fields: inspectFields,
	}
	return inspectItem
}

func toItemThinModel(item vault.Item) *ItemThinModel {
	return &ItemThinModel{
		Name:   item.Name(),
		Fields: item.FieldSize(),
	}
}

func toItemThinModels(profile vault.Profile) []*ItemThinModel {
	items := make([]*ItemThinModel, 0)
	for _, item := range profile.Items() {
		items = append(items, toItemThinModel(item))
	}
	return items
}

func toProfileInspectModel(profile vault.Profile) *ProfileInspectModel {
	inspectItems := make([]*ItemThinModel, 0)
	for _, item := range profile.Items() {
		inspectItems = append(inspectItems, toItemThinModel(item))
	}
	return &ProfileInspectModel{
		Name:  profile.Name(),
		Items: inspectItems,
	}
}

func toProfileNameOnlyModel(profileName string) *ProfileNameOnlyModel {
	return &ProfileNameOnlyModel{
		Name: profileName,
	}
}

func toProfileNameOnlyModels(profileNames []string) []*ProfileNameOnlyModel {
	profiles := make([]*ProfileNameOnlyModel, 0)
	for _, foundProfile := range profileNames {
		profiles = append(profiles, toProfileNameOnlyModel(foundProfile))
	}
	return profiles
}

func toVaultThinModel(vaultSetting model.VaultSetting) *VaultThinModel {
	return &VaultThinModel{
		Alias:      vaultSetting.Alias(),
		Identifier: vaultSetting.Identifier(),
	}
}

func toVaultThinModels(vaultSettings []model.VaultSetting) []*VaultThinModel {
	vaults := make([]*VaultThinModel, 0)
	for _, vaultSetting := range vaultSettings {
		vaults = append(vaults, toVaultThinModel(vaultSetting))
	}
	return vaults
}
