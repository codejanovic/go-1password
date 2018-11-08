package repository

import (
	"fmt"

	"github.com/codejanovic/gordon/environment"
	"github.com/codejanovic/gordon/fatal"
	"github.com/codejanovic/gordon/io"
	"github.com/codejanovic/gordon/model"
	"gopkg.in/yaml.v2"
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
		fatal.Crash(fmt.Errorf("We encountered a problem while persisting settings file"), "This looks like a programming error")
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
		fatal.Crash(err, "Error reading settings file")
	}
	err = yaml.Unmarshal(content, &settings)
	if err != nil {
		fatal.Crash(err, "Error interpreting settings file")
	}
	return &settings
}

func (s *yamlSettingsRepository) init() {
	settingsFile := io.NewFileByPath(environment.Environment.SettingsFile)
	if !settingsFile.Exists() {
		err := settingsFile.Create()
		if err != nil {
			fatal.Crash(err, "Unable to create required settings file")
		}
	}
}
