package repository

import (
	environment "github.com/codejanovic/go-1password/environment"
	io "github.com/codejanovic/go-1password/io"
	model "github.com/codejanovic/go-1password/model"
	throw "github.com/codejanovic/go-1password/throw"
	yaml "gopkg.in/yaml.v2"
)

type settingsRepositoryYaml struct {
}

var singleton SettingsRepository

func init() {
	singleton = &settingsRepositoryYaml{}
}

// NewSettingsRepositoryYaml constructor
func NewSettingsRepositoryYaml() SettingsRepository {
	return singleton
}

func (s *settingsRepositoryYaml) Fetch() model.Settings {
	s.init()
	return s.read()
}

func (s *settingsRepositoryYaml) Store(settings model.Settings) {
	settingsFile := io.NewFileByPath(environment.Environment.SettingsFile)
	data, err := yaml.Marshal(&settings)
	if err != nil {
		panic(err)
	}
	err = settingsFile.Write(data)
	if err != nil {
		panic(err)
	}
}

func (s *settingsRepositoryYaml) read() model.Settings {
	settingsFile := io.NewFileByPath(environment.Environment.SettingsFile)
	var settings model.SettingsYaml
	content, err := settingsFile.AsBytes()
	if err != nil {
		throw.Throw(err, "Error reading settings file")
	}
	err = yaml.Unmarshal(content, &settings)
	if err != nil {
		throw.Throw(err, "Error interpreting settings file")
	}
	return &settings
}

func (s *settingsRepositoryYaml) init() {
	settingsFile := io.NewFileByPath(environment.Environment.SettingsFile)
	if !settingsFile.Exists() {
		settingsFile.Create()
		s.Store(model.NewSettingsYaml())
	}
}
