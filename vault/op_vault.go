package vault

import (
	"errors"

	throw "github.com/codejanovic/go-1password/throw"
	"github.com/vinc3m1/opvault"
)

// OpVault struct
type opVault struct {
	vault *opvault.Vault
}

// NewOpVault constructor
func NewOpVault(path string) Vault {
	openedVault, err := opvault.Open(path)
	if err != nil {
		throw.Throw(err, "Unable to open opvault at "+path)
	}

	return &opVault{
		vault: openedVault,
	}
}

// TryOpen method
func (v *opVault) TryOpenProfile(name string, secret string) bool {
	_, err := v.OpenProfile(name, secret)
	return err == nil
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

	return profile, nil
}
