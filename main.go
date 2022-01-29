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
		Name:      "dbmagritte",
		Usage:     "DBMagritte is a tool for performing database migrations.",
		Version:   version,
		Copyright: "MIT",
		Authors: []*cli.Author{
			{
				Name:  "Austin Poor",
				Email: "code@austinpoor.com",
			},
		},
		EnableBashCompletion:   true,
		UseShortOptionHandling: true,
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
