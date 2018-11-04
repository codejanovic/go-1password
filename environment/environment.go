package environment

import (
	"path"

	homedir "github.com/mitchellh/go-homedir"
)

// Environment struct
type environment struct {
	UserDirectory string
	AppDirectory  string
	SettingsFile  string
}

// Environment singleton
var Environment *environment

// inits environment
func init() {
	userDir, err := homedir.Dir()
	if err != nil {
		panic("Unable to initialize environment, something went really wrong...")
	}

	appDir := path.Join(userDir, ".go-1password")

	env := &environment{
		UserDirectory: userDir,
		AppDirectory:  appDir,
		SettingsFile:  path.Join(appDir, "settings.yaml"),
	}
	Environment = env
}
