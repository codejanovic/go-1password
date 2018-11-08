package interaction

import (
	"fmt"
	"github.com/tcnksm/go-input"
	"os"
)

type consoleInteraction struct {
	ui *input.UI
}

var consoleInteractionSingleton *consoleInteraction

func init() {
	consoleInteractionSingleton = &consoleInteraction{
		ui: &input.UI{
			Writer: os.Stdout,
			Reader: os.Stdin,
		},
	}
}

func NewConsoleInteraction() Interaction {
	return consoleInteractionSingleton
}

func (c *consoleInteraction) AskForConfirmation(question string) (string, error) {
	query := question + " [Y/n]"
	answer, err := c.ui.Ask(query, &input.Options{
		Required: true,
		// Validate input
		ValidateFunc: func(s string) error {
			if s != "Y" && s != "n" {
				return fmt.Errorf("input must be Y or n")
			}

			return nil
		},
	})
	return answer, err
}

func (c *consoleInteraction) AskForVaultPassword() (string, error) {
	return c.AskForPassword("Please provide your vault secret first")
}

func (c *consoleInteraction) AskForPassword(command string) (string, error) {
	secret, err := c.ui.Ask(command, &input.Options{
		Required:    true,
		Mask:        true,
		MaskDefault: true,
	})

	return secret, err
}
