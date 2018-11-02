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
							output.WriteString(found.Alias)
						}
						log.Println(output.String())
						return nil
					},
				},
				{
					Name:  "add",
					Usage: "add {path/to/vault} {unique alias}",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
				{
					Name:  "signin",
					Usage: "signin {vault name} {vault password}",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
			},
		},
		{
			Name:  "profile",
			Usage: "profile optioms",
			Action: func(c *cli.Context) error {
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "list all profiles",
					Action: func(c *cli.Context) error {
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
