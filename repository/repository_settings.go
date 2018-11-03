package repository

import (
	model "github.com/codejanovic/go-1password/model"
)

// SettingsRepository interface
type SettingsRepository interface {
	Fetch() model.Settings
	Store(settings model.Settings)
}

var settingsRepositorySingleton SettingsRepository

func init() {
	settingsRepositorySingleton = &yamlSettingsRepository{}
}

// NewSettingsRepository constructor
func NewSettingsRepository() SettingsRepository {
	return settingsRepositorySingleton
}
