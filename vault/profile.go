package vault

import (
	fatal "github.com/codejanovic/go-1password/fatal"
	"github.com/vinc3m1/opvault"
)

// Profile interface
type Profile interface {
	Name() string
	ItemsSize() int
	Items() []Item
}

type opVaultProfile struct {
	profileInternal *opvault.Profile
	itemsInternal   []Item
}

func newOpVaultProfile(originalProfile *opvault.Profile) Profile {
	originalItems, err := originalProfile.Items()
	if err != nil {
		fatal.Crash(err, "Unable to fetch items from profile")
	}

	items := make([]Item, 0)
	for _, originalItem := range originalItems {
		items = append(items, newOpVaultItem(originalItem))
	}

	return &opVaultProfile{
		profileInternal: originalProfile,
		itemsInternal:   items,
	}
}

func (v *opVaultProfile) Name() string {
	return v.profileInternal.ProfileName()
}

func (v *opVaultProfile) ItemsSize() int {
	return len(v.itemsInternal)
}

func (v *opVaultProfile) Items() []Item {
	return v.itemsInternal
}
