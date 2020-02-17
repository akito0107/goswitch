package main

import (
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
					return nil
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

	app.Run(os.Args)
}
