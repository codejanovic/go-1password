package model

// VaultSetting interface for saving vault specific settings
type VaultSetting interface {
	Identifier() string
	Alias() string
	WithAlias(alias string)
	Path() string
	Profile() string
	WithProfile(profile string)
	IsEqualTo(identifierOrAlias string) bool
}
