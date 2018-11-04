package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	usecase "github.com/codejanovic/go-1password/usecase"
	"github.com/urfave/cli"
)

var (
	version = "v0.1.0"
	appName = "go-1password"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "cli for interacting with local opvault vaults"
	app.Version = version
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
						if !response.HasVaults() {
							return errors.New("no configured vaults found")
						}

						data, err := json.Marshal(response.Vaults)
						if err != nil {
							return err
						}
						log.Println(string(data))
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
						data, err := json.Marshal(response.Vault)
						if err != nil {
							return err
						}
						log.Println("successfully added vault " + string(data))
						return nil
					},
				},
				{
					Name:      "remove",
					Usage:     "remove a vault from your configuration",
					ArgsUsage: "$0{vault identifier|alias}",
					Action: func(c *cli.Context) error {
						err := usecase.NewRemoveVaultUsecase().Execute(&usecase.RemoveVaultRequest{
							VaultAliasOrIdentifier: c.Args().Get(0),
						})
						if err != nil {
							return err
						}
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
					Name:  "list",
					Usage: "list all profiles within a vault",
					Action: func(c *cli.Context) error {
						response, err := usecase.NewListProfileUsecase().Execute()
						if err != nil {
							return err
						}
						data, err := json.Marshal(response.Profiles)
						if err != nil {
							return err
						}
						log.Println(string(data))
						return nil
					},
				},
				{
					Name:  "inspect",
					Usage: "inspect profile",
					Action: func(c *cli.Context) error {
						response, err := usecase.NewInspectProfileUsecase().Execute()
						if err != nil {
							return err
						}
						data, err := json.Marshal(response.Profile)
						if err != nil {
							return err
						}
						log.Println(string(data))
						return nil
					},
				},
			},
		},
		{
			Name:  "item",
			Usage: "item options",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "list all items within a profile",
					Action: func(c *cli.Context) error {
						response, err := usecase.NewListItemsUsecase().Execute()
						if err != nil {
							return err
						}
						data, err := json.Marshal(response.Items)
						if err != nil {
							return err
						}
						log.Println(string(data))
						return nil
					},
				},
				{
					Name:      "inspect",
					Usage:     "inspect item",
					ArgsUsage: "$0{item name}",
					Action: func(c *cli.Context) error {
						response, err := usecase.NewInspectItemUsecase().Execute(&usecase.InspectItemRequest{
							ItemName: c.Args().Get(0),
						})
						if err != nil {
							return err
						}
						data, err := json.Marshal(response.Item)
						if err != nil {
							return err
						}
						log.Println(string(data))
						return nil
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
