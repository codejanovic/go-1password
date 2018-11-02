package repository

import (
	model "github.com/codejanovic/go-1password/model"
)
type SettingsRepository interface {
	Fetch() model.Settings
	Store(settings model.Settings)
}