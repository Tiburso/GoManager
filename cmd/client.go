package cmd

import (
	"github.com/urfave/cli/v2"
)

var ClientCommand *cli.Command = &cli.Command{
	Name:        "client",
	Aliases:     []string{"c"},
	Usage:       "client commands",
	Description: "client commands",
	Subcommands: []*cli.Command{
		CompanyCommand,
		ApplicationCommand,
	},
}
