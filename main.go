package main

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	usecase "github.com/codejanovic/go-1password/usecase"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-1password"
	app.Usage = "cli for interacting with local opvault vaults"
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "codejanovic",
			Email: "codejanovic@gmail.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "vault",
			Usage: "",
			Subcommands: []cli.Command{
				{
					Name:      "list",
					Usage:     "list all configured vaults",
					ArgsUsage: "no arguments",
					Action: func(c *cli.Context) error {
						response := usecase.NewListVaultUsecase().Execute()
						var output strings.Builder
						if !response.HasVaults() {
							return errors.New("there are no vaults configured")
						}

						for i, found := range response.Found {
							if i > 0 {
								output.WriteString(", ")
							}
							output.WriteString("[ ")
							output.WriteString("id=" + found.Identifier)
							output.WriteString(", alias=" + found.Alias)
							output.WriteString(" ]")
						}
						log.Println(output.String())
						return nil
					},
				},
				{
					Name:      "add",
					Usage:     "add a vault to your configuration",
					ArgsUsage: "$0{vault path}, $1{vault alias}",
					Action: func(c *cli.Context) error {
						response, err := usecase.NewAddVaultUsecase().Execute(&usecase.AddVaultRequest{
							VaultPath:  c.Args().Get(0),
							VaultAlias: c.Args().Get(1),
						})
						if err != nil {
							return err
						}
						log.Printf("successfully added vault '%s' (%s)", response.Added.Alias(), response.Added.Path())
						return nil
					},
				},
				{
					Name:      "remove",
					Usage:     "remove a vault from your configuration",
					ArgsUsage: "$1{vault identifier|alias}",
					Action: func(c *cli.Context) error {
						usecase.NewRemoveVaultUsecase().Execute(&usecase.RemoveVaultRequest{
							VaultAliasOrIdentifier: c.Args().Get(0),
						})
						log.Printf("successfully removed vault '%s'", c.Args().Get(0))
						return nil
					},
				},
				{
					Name:      "signin",
					Usage:     "signin {vault name} {vault password}",
					ArgsUsage: "$0{vault identifier|alias}, $1{vault profile}, $2{vault password}",
					Action: func(c *cli.Context) error {
						err := usecase.NewSignInVaultUsecase().Execute(&usecase.SignInVaultRequest{
							VaultAliasOrIdentifier: c.Args().Get(0),
							VaultProfile:           c.Args().Get(1),
							VaultSecret:            c.Args().Get(2),
						})
						if err != nil {
							return err
						}
						log.Printf("successfully signed into vault '%s' with profile '%s'", c.Args().Get(0), c.Args().Get(1))
						return nil
					},
				},
			},
		},
		{
			Name:  "profile",
			Usage: "profile options",
			Subcommands: []cli.Command{
				{
					Name:      "list",
					Usage:     "list all profiles within a vault",
					ArgsUsage: "$0{vault identifier|alias}",
					Action: func(c *cli.Context) error {
						response, err := usecase.NewListProfileUsecase().Execute(&usecase.ListProfileRequest{
							VaultAliasOrIdentifier: c.Args().Get(0),
						})
						if err != nil {
							return err
						}
						var output strings.Builder
						for i, profile := range response.Found {
							if i > 0 {
								output.WriteString(", ")
							}
							output.WriteString(profile)
						}
						log.Println("profiles found: " + output.String())
						return nil
					},
				},
				{
					Name:  "signin",
					Usage: "signin {profile name}",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
				{
					Name:  "search",
					Usage: "search for profiles",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
				{
					Name:  "inspect",
					Usage: "inspect profile",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
				{
					Name:  "items",
					Usage: "items",
					Subcommands: []cli.Command{
						{
							Name:  "show",
							Usage: "show specific item",
							Action: func(c *cli.Context) error {
								return nil
							},
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
