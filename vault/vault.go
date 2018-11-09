package vault

import (
	"errors"

	"github.com/codejanovic/gordon/fatal"
	"github.com/vinc3m1/opvault"
)

// Vault interface
type Vault interface {
	TryOpenProfile(name string, secret string) bool
	OpenProfile(name string, secret string) (Profile, error)
	OpenDefaultProfile(secret string) (Profile, error)
	Profiles() ([]string, error)
	HasDefaultProfile() bool
}

type opVault struct {
	vault   *opvault.Vault
	profile string
}

// NewOpVault constructor
func NewOpVault(path string, defaultProfile string) Vault {
	openedVault, err := opvault.Open(path)
	if err != nil {
		fatal.Crash(err, "unable to open opvault at "+path)
	}

	return &opVault{
		vault:   openedVault,
		profile: defaultProfile,
	}
}

// HasDefaultProfile
func (v *opVault) HasDefaultProfile() bool {
	return v.profile != ""
}

// TryOpen method
func (v *opVault) TryOpenProfile(name string, secret string) bool {
	_, err := v.OpenProfile(name, secret)
	return err == nil
}

// Profiles get all found profile names
func (v *opVault) Profiles() ([]string, error) {
	names, err := v.vault.ProfileNames()
	if err != nil {
		return nil, errors.New("Unable to fetch profiles from vault. Reason: " + err.Error())
	}
	return names, nil
}

// OpenDefaultProfile method
func (v *opVault) OpenDefaultProfile(secret string) (Profile, error) {
	return v.OpenProfile(v.profile, secret)
}

// OpenProfile method
func (v *opVault) OpenProfile(name string, secret string) (Profile, error) {
	profile, err := v.vault.Profile(name)
	if err != nil {
		return nil, errors.New("Unable to find Profile. Reason: " + err.Error())
	}
	err = profile.Unlock(secret)
	if err != nil {
		return nil, errors.New("Unable to unlock Profile. Reason: " + err.Error())
	}

	return newOpVaultProfile(profile), nil
}
