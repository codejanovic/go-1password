package model

import (
	"path/filepath"
	"strings"
)

// VaultSettingYaml struct
type VaultSettingYaml struct {
	IdentifierInternal string `yaml:"identifier"`
	PathInternal       string `yaml:"path"`
	ActiveInternal     bool   `yaml:"active"`
	ProfileInternal    string `yaml:"profile"`
	AliasInternal      string `yaml:"alias"`
}

// Path returning path of configured vault
func (v *VaultSettingYaml) Path() string {
	absolutePath, err := filepath.Abs(v.PathInternal)
	if err != nil {
		panic(err)
	}
	return absolutePath
}

// Identifier returning the vault identifier
func (v *VaultSettingYaml) Identifier() string {
	return v.IdentifierInternal
}

// Profile returning the vault identifier
func (v *VaultSettingYaml) Profile() string {
	return v.ProfileInternal
}

// WithProfile returning the vault identifier
func (v *VaultSettingYaml) WithProfile(profileName string) {
	v.ProfileInternal = profileName
}

// Alias returning the vault identifier
func (v *VaultSettingYaml) Alias() string {
	return v.AliasInternal
}

// WithAlias returning the vault identifier
func (v *VaultSettingYaml) WithAlias(alias string) {
	v.AliasInternal = alias
}

// IsEqualTo checks equality
func (v *VaultSettingYaml) IsEqualTo(identifierOrAlias string) bool {
	return strings.EqualFold(v.IdentifierInternal, identifierOrAlias) || strings.EqualFold(v.AliasInternal, identifierOrAlias)
}
