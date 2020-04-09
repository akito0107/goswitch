package main

import (
	"log"
	"os"

	"github.com/akito0107/goswitch/internal"
	"github.com/urfave/cli/v2"
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
				Action: func(c *cli.Context) error {
					return internal.LSRemote(c.Context)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
