package environment

import (
	"path"

	throw "github.com/codejanovic/go-1password/throw"
	"github.com/docker/docker-credential-helpers/credentials"
	homedir "github.com/mitchellh/go-homedir"
)

// Environment struct
type environment struct {
	UserDirectory string
	AppDirectory  string
	SettingsFile  string
	nativeStore   credentials.Helper
}

func (e *environment) AddCredentials(credentials *credentials.Credentials) {
	err := e.nativeStore.Add(credentials)
	if err != nil {
		throw.Throw(err, "Did you deny credentials access to go-1password?")
	}
}

func (e *environment) GetCredentials(identifier string) (string, string) {
	user, secret, err := e.nativeStore.Get(identifier)
	if err != nil {
		throw.Throw(err, "Did you deny credentials access to go-1password?")
	}
	return user, secret
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
