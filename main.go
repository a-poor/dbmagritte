package main

import (
	"fmt"
	"log"
	"os"
	"path"

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
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "path",
				Usage:       "`PATH` to the project root.",
				DefaultText: `$PWD`,
				Value:       ".",
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "init",
				Usage:       "Initialize the current project for migrations.",
				Description: cmdInitDesc,
				Category:    "config",
				Action: func(c *cli.Context) error {
					// Get the path to the project root.
					p := c.String("path")

					// Is that a valid location?
					if !isAtGitRoot(p) {
						return cli.Exit(
							"Not at a valid project root.",
							1,
						)
					}

					// Create the config file.
					conf := newGlobalConfig()
					cf := path.Join(p, defaultGlobalConfigPath)
					err := conf.write(cf)
					if err != nil {
						log.Printf("Error writing config file: %q\n", err)
						return cli.Exit(
							"Failed to write config file.",
							1,
						)
					}
					fmt.Printf("Wrote config file to: %s\n", cf)

					// Create the migrations directory.
					md := path.Join(p, "migrations")
					err = os.MkdirAll(md, 0755)
					if err != nil {
						log.Printf("Failed to create migrations directory: %q\n", err)
						return cli.Exit(
							"Unable to create migrations directory",
							1,
						)
					}
					fmt.Printf("Created the migrations directory: %s\n", md)

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
					// Get the path to the project root.
					p := c.String("path")

					// Is that a valid location?
					if !isAtGitRoot(p) {
						return cli.Exit(
							"Not at a valid project root.",
							1,
						)
					}

					// Get the current git hash
					hash, err := getGitHash(p)
					if err != nil {
						return cli.Exit(
							"Failed to get the current git hash.",
							1,
						)
					}

					// Check if this migration already exists
					md := path.Join(p, "migrations", hash)
					err = os.Mkdir(md, 0755)
					if err != nil {
						return cli.Exit(
							"Failed to create migration directory.",
							1,
						)
					}
					fmt.Printf("Created the migration directory: %s\n", md)

					// Create the migration files
					mup := path.Join(md, "up.sql")
					os.OpenFile(
						mup,
						os.O_RDONLY|os.O_CREATE,
						0666,
					)
					mdown := path.Join(md, "down.sql")
					os.OpenFile(
						mdown,
						os.O_RDONLY|os.O_CREATE,
						0666,
					)

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
				Category:    "other",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:        "validate",
				Aliases:     []string{"v"},
				Usage:       "Validate the current project state.",
				Description: cmdValidateDesc,
				Category:    "other",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
		// Action: func(c *cli.Context) error {
		// 	h, err := getGitHash()
		// 	if err == ErrNotAtGitRoot {
		// 		return cli.Exit("No git repository found at current directory.", 1)
		// 	}
		// 	if err != nil {
		// 		return cli.Exit("Unable to read git repo.", 1)
		// 	}
		// 	fmt.Println("Git hash:", h)

		// 	return nil
		// },
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
