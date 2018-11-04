package repository

import "github.com/docker/docker-credential-helpers/wincred"

func init() {
	dockerCredentialsRepositorySingleton = &dockerCredentialsRepository{
		nativeStore: wincred.Wincred{},
	}
	credentialsRepositorySingleton = dockerCredentialsRepositorySingleton
}
