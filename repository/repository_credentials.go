package repository

// CredentialsRepository interface
type CredentialsRepository interface {
	Fetch(identifier string) string
	Remove(identifier string)
	Store(identifier string, secret string)
}

var credentialsRepositorySingleton CredentialsRepository

// NewCredentialsRepository constructor
func NewCredentialsRepository() CredentialsRepository {
	return credentialsRepositorySingleton
}
