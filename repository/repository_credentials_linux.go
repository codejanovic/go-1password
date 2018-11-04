package repository

import (
	"github.com/codejanovic/go-1password/fatal"
	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/docker/docker-credential-helpers/secretservice"
)

// linuxCredentialsRepository struct
type linuxCredentialsRepository struct {
	nativeStore credentials.Helper
}

var linuxCredentialsRepositorySingleton *linuxCredentialsRepository

func init() {
	linuxCredentialsRepositorySingleton = &linuxCredentialsRepository{
		nativeStore: secretservice.Secretservice{},
	}
	credentialsRepositorySingleton = linuxCredentialsRepositorySingleton
}

// Store credentials
func (r *linuxCredentialsRepository) Store(identifier string, secret string) {
	err := r.nativeStore.Add(&credentials.Credentials{
		ServerURL: identifier,
		Secret:    secret,
	})

	if err != nil {
		fatal.Crash(err, "Unable to access credentials store")
	}
}

// Fetch credentials
func (r *linuxCredentialsRepository) Fetch(identifier string) (string, bool) {
	_, secret, err := r.nativeStore.Get(identifier)
	if err != nil {
		return "", false
	}
	return secret, true
}

// Remove credentials
func (r *linuxCredentialsRepository) Remove(identifier string) {
	err := r.nativeStore.Delete(identifier)
	if err != nil {
		fatal.Crash(err, "Unable to access credentials store")
	}
}
