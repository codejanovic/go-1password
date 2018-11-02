package model

import (
	"path/filepath"
	"strings"
)

// VaultSettingYaml struct
type VaultSettingYaml struct {
	identifier string `yaml:"identifier"`
	path       string `yaml:"path"`
	active     bool   `yaml:"active"`
	profile    string `yaml:"profile"`
	alias      string `yaml:"alias"`
}

// Path returning path of configured vault
func (v *VaultSettingYaml) Path() string {
	absolutePath, err := filepath.Abs(v.path)
	if err != nil {
		panic(err)
	}
	return absolutePath
}

// Identifier returning the vault identifier
func (v *VaultSettingYaml) Identifier() string {
	return v.identifier
}

// Profile returning the vault identifier
func (v *VaultSettingYaml) Profile() string {
	return v.profile
}

// WithProfile returning the vault identifier
func (v *VaultSettingYaml) WithProfile(profileName string) {
	v.profile = profileName
}

// Alias returning the vault identifier
func (v *VaultSettingYaml) Alias() string {
	return v.alias
}

// WithAlias returning the vault identifier
func (v VaultSettingYaml) WithAlias(alias string) {
	v.alias = alias
}

// IsEqualTo checks equality
func (v *VaultSettingYaml) IsEqualTo(identifierOrAlias string) bool {
	return strings.EqualFold(v.identifier, identifierOrAlias) || strings.EqualFold(v.alias, identifierOrAlias)
}
