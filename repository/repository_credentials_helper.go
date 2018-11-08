package repository

import (
	"github.com/codejanovic/gordon/fatal"
	"github.com/docker/docker-credential-helpers/credentials"
)

// dockerCredentialsRepository struct
type dockerCredentialsRepository struct {
	nativeStore credentials.Helper
}

var dockerCredentialsRepositorySingleton *dockerCredentialsRepository

// Store credentials
func (r *dockerCredentialsRepository) Store(identifier string, secret string) {
	err := r.nativeStore.Add(&credentials.Credentials{
		ServerURL: identifier,
		Secret:    secret,
	})

	if err != nil {
		fatal.Crash(err, "Unable to access credentials store")
	}
}

// Fetch credentials
func (r *dockerCredentialsRepository) Fetch(identifier string) (string, bool) {
	_, secret, err := r.nativeStore.Get(identifier)
	if err != nil {
		return "", false
	}
	return secret, true
}

// Remove credentials
func (r *dockerCredentialsRepository) Remove(identifier string) {
	err := r.nativeStore.Delete(identifier)
	if err != nil {
		fatal.Crash(err, "Unable to access credentials store")
	}
}
