package repository

import (
	throw "github.com/codejanovic/go-1password/throw"
	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/docker/docker-credential-helpers/secretservice"
)

// LinuxCredentialsRepository struct
type LinuxCredentialsRepository struct {
	nativeStore credentials.Helper
}

var linuxCredentialsRepositorySingleton *LinuxCredentialsRepository

func init() {
	linuxCredentialsRepositorySingleton = &LinuxCredentialsRepository{
		nativeStore: secretservice.Secretservice{},
	}
	credentialsRepositorySingleton = linuxCredentialsRepositorySingleton
}

// Store credentials
func (r *LinuxCredentialsRepository) Store(identifier string, secret string) {
	err := r.nativeStore.Add(&credentials.Credentials{
		ServerURL: identifier,
		Secret:    secret,
	})

	if err != nil {
		throw.Throw(err, "Unable to access credentials store")
	}
}

// Fetch credentials
func (r *LinuxCredentialsRepository) Fetch(identifier string) string {
	_, secret, err := r.nativeStore.Get(identifier)
	if err != nil {
		throw.Throw(err, "Unable to access credentials store")
	}
	return secret
}

// Remove credentials
func (r *LinuxCredentialsRepository) Remove(identifier string) {
	err := r.nativeStore.Delete(identifier)
	if err != nil {
		throw.Throw(err, "Unable to access credentials store")
	}
}
