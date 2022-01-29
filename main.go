package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// version is the current version of the CLI.
const version = "0.1.0"

func main() {
	app := &cli.App{
		Name:        "dbmagritte",
		Usage:       "A tool for performing database migrations.",
		Description: appDesc,
		Version:     version,
		Copyright:   "Copyright (c) 2022 Austin Poor",
		Authors: []*cli.Author{
			{
				Name:  "Austin Poor",
				Email: "code@austinpoor.com",
			},
		},
		EnableBashCompletion:   true,
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:        "init",
				Usage:       "Initialize the current git repo for migrations.",
				Description: cmdInitDesc,
				Category:    "config",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:        "new",
				Aliases:     []string{"n"},
				Usage:       "Create a new migration.",
				Description: cmdNewDesc,
				Category:    "config",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:        "info",
				Aliases:     []string{"i"},
				Usage:       "Get info about the current migration state.",
				Description: cmdInfoDesc,
				Category:    "config",
				Action: func(c *cli.Context) error {
					return nil
				},
			},

			{
				Name:        "up",
				Aliases:     []string{"u"},
				Usage:       "Move forward in the migration tree.",
				Description: cmdUpDesc,
				Category:    "migrate",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:        "down",
				Aliases:     []string{"d"},
				Usage:       "Move backward in the migration tree.",
				Description: cmdDownDesc,
				Category:    "migrate",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:        "fast-forward",
				Aliases:     []string{"ff"},
				Usage:       "Move to the newest migration.",
				Description: cmdFastForwardDesc,
				Category:    "migrate",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:        "reset",
				Usage:       "Rollback all migrations.",
				Description: cmdResetDesc,
				Category:    "migrate",
				Action: func(c *cli.Context) error {
					return nil
				},
			},

			{
				Name:        "check",
				Usage:       "Check the connection to the database.",
				Description: cmdCheckDesc,
				Category:    "db",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:        "validate",
				Aliases:     []string{"v"},
				Usage:       "Validate the current migration state and directory structure.",
				Description: cmdValidateDesc,
				Category:    "db",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			h, err := getGitHash()
			if err == ErrNotAtGitRoot {
				fmt.Println("Error: No git repository found at current directory.")
				return nil
			}
			if err != nil {
				fmt.Println("Error: Unable to read git repo.")
				return err
			}
			fmt.Println("Git hash:", h)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
