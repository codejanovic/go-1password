package fatal

import (
	"strings"
)

// Crash indicates a error in the program and prints out some help
func Crash(err error, hint string) {
	var customError strings.Builder
	customError.WriteString("Oops, something was missing or unexpected..\n")
	customError.WriteString("Reason was: " + err.Error() + "\n")
	customError.WriteString("Hint is: " + hint + "\n")
	customError.WriteString("If the reason and hint are unclear, please dont hestitate to file a ticket!\n")
	customError.WriteString("Github: https://github.com/codejanovic/gordon\n")
	panic(customError.String())
}
