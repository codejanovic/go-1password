package repository

import "github.com/docker/docker-credential-helpers/secretservice"

func init() {
	dockerCredentialsRepositorySingleton = &dockerCredentialsRepository{
		nativeStore: secretservice.Secretservice{},
	}
	credentialsRepositorySingleton = dockerCredentialsRepositorySingleton
}
