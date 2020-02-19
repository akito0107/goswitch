package main

import (
	"log"
	"os"

	"github.com/akito0107/switchgo/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "use",
				Action: func(c *cli.Context) error {
					v := c.Args().Get(0)
					return internal.Use(c.Context, v)
				},
			},
			{
				Name: "ls",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name: "ls-remote",
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
