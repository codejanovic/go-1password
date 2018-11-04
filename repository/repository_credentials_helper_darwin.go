package repository

import "github.com/docker/docker-credential-helpers/osxkeychain"

func init() {
	dockerCredentialsRepositorySingleton = &dockerCredentialsRepository{
		nativeStore: osxkeychain.Osxkeychain{},
	}
	credentialsRepositorySingleton = dockerCredentialsRepositorySingleton
}
