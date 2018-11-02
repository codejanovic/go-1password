package environment

import (
	"github.com/docker/docker-credential-helpers/wincred"
)

// inits environment
func init() {
	nativeStore := wincred.Wincred{}
	Environment.nativeStore = nativeStore
}
