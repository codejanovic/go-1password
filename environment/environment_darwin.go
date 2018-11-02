package environment

import (
	"github.com/docker/docker-credential-helpers/osxkeychain"
)

// inits environment
func init() {
	nativeStore := osxkeychain.Osxkeychain{}
	Environment.nativeStore = nativeStore
}
