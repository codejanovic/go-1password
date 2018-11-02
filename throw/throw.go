package fatal

import (
	"strings"
)

// Throw s with message and some useful Information for the user
func Throw(err error, hint string) {
	var customError strings.Builder
	customError.WriteString("Oops, something was missing or unexpected..\n")
	customError.WriteString("Reason was: " + err.Error() + "\n")
	customError.WriteString("Hint is: " + hint + "\n")
	customError.WriteString("If the reason and hint are unclear, please dont hestitate to file a ticket!\n")
	customError.WriteString("Github: https://github.com/codejanovic/go-1password\n")
	panic(customError.String())
}
