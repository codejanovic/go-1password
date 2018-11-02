package model

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

// SettingsYaml in yaml
type SettingsYaml struct {
	vaults []*VaultSettingYaml `yaml:"vaults"`
}

// NewSettingsYaml constructor
func NewSettingsYaml() Settings {
	return &SettingsYaml{
		vaults: make([]*VaultSettingYaml, 0),
	}
}

// Vaults returning configured vaults
func (s *SettingsYaml) Vaults() []VaultSetting {
	if s.vaults == nil {
		return make([]VaultSetting, 0)
	}
	vaultSettings := make([]VaultSetting, len(s.vaults))
	for i, v := range s.vaults {
		vaultSettings[i] = v
	}
	return vaultSettings
}

// Active returning configured vaults
func (s *SettingsYaml) Active() (VaultSetting, error) {
	for _, v := range s.vaults {
		if v.active {
			return v, nil
		}
	}

	return nil, errors.New("unable to find active vault. Did you forget to add a vault in advance?")
}

// Add and return vault
func (s *SettingsYaml) Add(path string, alias string) (VaultSetting, error) {
	if s.vaults == nil {
		s.vaults = make([]*VaultSettingYaml, 0)
	}

	_, err := s.Find(alias)
	if err == nil {
		return nil, errors.New("a vault with alias " + alias + " already exists")
	}

	vaultToAdd := &VaultSettingYaml{
		active:     false,
		path:       path,
		alias:      alias,
		identifier: uuid.Must(uuid.NewV4()).String(),
	}

	s.vaults = append(s.vaults, vaultToAdd)
	s.Activate(vaultToAdd.identifier)
	return vaultToAdd, nil
}

// Activate a vault
func (s *SettingsYaml) Activate(identifierOrAlias string) VaultSetting {
	var activatedVault VaultSetting
	for _, v := range s.vaults {
		if v.IsEqualTo(identifierOrAlias) {
			v.active = true
			activatedVault = v
		} else {
			v.active = false
		}
	}
	return activatedVault
}

// Find a vault
func (s *SettingsYaml) Find(identifierOrAlias string) (VaultSetting, error) {
	for _, vault := range s.vaults {
		if vault.IsEqualTo(identifierOrAlias) {
			return vault, nil
		}
	}
	return nil, errors.New("unable to find vault with identifier or alias " + identifierOrAlias)
}

// Remove a vault
func (s *SettingsYaml) Remove(identifierOrAlias string) {

}
