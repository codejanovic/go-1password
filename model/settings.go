package model

// Settings interface
type Settings interface {
	Vaults() []VaultSetting
	Active() (VaultSetting, error)
	Add(path string, alias string) (VaultSetting, error)
	Remove(identifierOrAlias string)
	Activate(identifierOrAlias string) VaultSetting
	Find(identifierOrAlias string) (VaultSetting, error)
}
