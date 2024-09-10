package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/akito0107/goswitch/internal/v2"
)

func main() {
	app := &cli.App{
		Name:  "goswitch",
		Usage: "go version switching utility",
		Commands: []*cli.Command{
			{
				Name:  "use",
				Usage: "switch current go version",
				Action: func(c *cli.Context) error {
					v := c.Args().Get(0)
					return internal.Use(c.Context, v)
				},
			},
			{
				Name:  "ls-remote",
				Usage: "show all available versions",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:     "use-github",
						Aliases:  []string{"g"},
						Usage:    "use github tags",
						Required: false,
					},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("use-github") {
						log.Println("Using github releases is now enabled by default, so this option no longer has any effect.")
					}
					return internal.LSRemoteGH(c.Context)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
