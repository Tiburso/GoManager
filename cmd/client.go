package cmd

import (
	"github.com/urfave/cli/v2"
)

func ClientCommand() *cli.Command {
	return &cli.Command{
		Name:        "client",
		Aliases:     []string{"c"},
		Usage:       "client commands",
		Description: "client commands",
		Subcommands: []*cli.Command{
			CompanyCommand,
			ApplicationCommand,
		},
	}
}
