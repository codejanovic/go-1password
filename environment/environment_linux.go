package environment

import (
	"github.com/docker/docker-credential-helpers/secretservice"
)

// inits environment
func init() {
	nativeStore := secretservice.Secretservice{}
	Environment.nativeStore = nativeStore
}
