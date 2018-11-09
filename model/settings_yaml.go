package model

import (
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
)

// SettingsYaml in yaml
type SettingsYaml struct {
	VaultsInternal []*VaultSettingYaml `yaml:"vaults"`
}

// NewSettingsYaml constructor
func NewSettingsYaml() Settings {
	return &SettingsYaml{
		VaultsInternal: make([]*VaultSettingYaml, 0),
	}
}

// Vaults returning configured vaults
func (s *SettingsYaml) Vaults() []VaultSetting {
	if s.VaultsInternal == nil {
		return make([]VaultSetting, 0)
	}
	vaultSettings := make([]VaultSetting, len(s.VaultsInternal))
	for i, v := range s.VaultsInternal {
		vaultSettings[i] = v
	}
	return vaultSettings
}

// Active returning configured vaults
func (s *SettingsYaml) ActiveOrAlternative(altIdentifierOrAlias string) (VaultSetting, error) {
	var foundVaultSetting VaultSetting
	var err error
	if altIdentifierOrAlias == "" {
		foundVaultSetting, err = s.Active()
		if err != nil {
			return nil, fmt.Errorf("unable to find active vault. did you sign in into a default vault and profile? ")
		}
	} else {
		foundVaultSetting, err = s.Find(altIdentifierOrAlias)
		if err != nil {
			return nil, fmt.Errorf("unable to find vault %s. did you configure this vault in advance? ", altIdentifierOrAlias)
		}
	}
	return foundVaultSetting, nil
}

// Active returning configured vaults
func (s *SettingsYaml) Active() (VaultSetting, error) {
	for _, v := range s.VaultsInternal {
		if v.ActiveInternal {
			return v, nil
		}
	}

	return nil, errors.New("unable to find active vault. Did you forget to add a vault in advance?")
}

// Add and return vault
func (s *SettingsYaml) Add(vaultPath string, vaultAlias string) (VaultSetting, error) {
	if s.VaultsInternal == nil {
		s.VaultsInternal = make([]*VaultSettingYaml, 0)
	}

	_, err := s.Find(vaultAlias)
	if err == nil {
		return nil, errors.New("a vault with alias " + vaultAlias + " already exists")
	}

	vaultToAdd := &VaultSettingYaml{
		ActiveInternal:     false,
		PathInternal:       vaultPath,
		AliasInternal:      vaultAlias,
		IdentifierInternal: uuid.Must(uuid.NewV4()).String(),
	}

	s.VaultsInternal = append(s.VaultsInternal, vaultToAdd)
	s.Activate(vaultToAdd.IdentifierInternal)
	return vaultToAdd, nil
}

// Activate a vault
func (s *SettingsYaml) Activate(identifierOrAlias string) VaultSetting {
	var activatedVault VaultSetting
	for _, v := range s.VaultsInternal {
		if v.IsEqualTo(identifierOrAlias) {
			v.ActiveInternal = true
			activatedVault = v
		} else {
			v.ActiveInternal = false
		}
	}
	return activatedVault
}

// Find a vault
func (s *SettingsYaml) Find(identifierOrAlias string) (VaultSetting, error) {
	for _, vault := range s.VaultsInternal {
		if vault.IsEqualTo(identifierOrAlias) {
			return vault, nil
		}
	}
	return nil, errors.New("unable to find vault with identifier or alias " + identifierOrAlias)
}

// Remove a vault
func (s *SettingsYaml) Remove(identifierOrAlias string) {
	vaultToRemove := -1
	for i, vault := range s.VaultsInternal {
		if vault.IsEqualTo(identifierOrAlias) {
			vaultToRemove = i
			break
		}
	}

	if vaultToRemove != -1 {
		s.VaultsInternal = append(s.VaultsInternal[:vaultToRemove], s.VaultsInternal[vaultToRemove+1:]...)
	}
}
