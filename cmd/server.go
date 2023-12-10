package main

import (
	"github.com/Tiburso/GoManager/common"
	"github.com/Tiburso/GoManager/models/db"
	"github.com/Tiburso/GoManager/routers"
	"github.com/urfave/cli/v2"
)

func RunServer(cCtx *cli.Context) error {
	common.LoadEnv()

	// This sets up the db connection
	db.ConnectDatabase()

	// The server can only run when the db connection is established
	routers.RunServer()

	return nil
}

func ServerCommand() *cli.Command {
	return &cli.Command{
		Name:        "server",
		Aliases:     []string{"s"},
		Usage:       "server commands",
		Description: "server commands",
		Subcommands: []*cli.Command{
			{
				Name:        "run",
				Aliases:     []string{"r"},
				Usage:       "run server",
				Description: "run server",
				Action:      RunServer,
			},
		},
	}
}
