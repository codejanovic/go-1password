package repository

import (
	"fmt"

	environment "github.com/codejanovic/go-1password/environment"
	io "github.com/codejanovic/go-1password/io"
	model "github.com/codejanovic/go-1password/model"
	throw "github.com/codejanovic/go-1password/throw"
	yaml "gopkg.in/yaml.v2"
)

type yamlSettingsRepository struct {
}

func (s *yamlSettingsRepository) Fetch() model.Settings {
	s.init()
	return s.read()
}

func (s *yamlSettingsRepository) Store(settings model.Settings) {
	yamlSettings, ok := settings.(*model.SettingsYaml)
	if !ok {
		throw.Throw(fmt.Errorf("We encountered a problem while persisting settings file"), "This looks like a programming error")
	}

	settingsFile := io.NewFileByPath(environment.Environment.SettingsFile)
	data, err := yaml.Marshal(&yamlSettings)
	if err != nil {
		panic(err)
	}

	err = settingsFile.Write(data)
	if err != nil {
		panic(err)
	}
}

func (s *yamlSettingsRepository) read() model.Settings {
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

func (s *yamlSettingsRepository) init() {
	settingsFile := io.NewFileByPath(environment.Environment.SettingsFile)
	if !settingsFile.Exists() {
		settingsFile.Create()
	}
}
