package cmd

import (
	"github.com/Tiburso/GoManager/common"
	"github.com/Tiburso/GoManager/models/db"
	"github.com/Tiburso/GoManager/routers"
	"github.com/urfave/cli/v2"
)

func RunServer(cCtx *cli.Context) error {
	common.LoadEnv()

	port := cCtx.String("port")

	// This sets up the db connection
	if err := db.ConnectDatabase(); err != nil {
		return err
	}

	// The server can only run when the db connection is established
	routers.RunServer(port)

	return nil
}

var ServerCommand *cli.Command = &cli.Command{
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
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "port",
					Aliases:     []string{"p"},
					Usage:       "server port",
					Required:    false,
					DefaultText: "8080",
					Value:       "8080",
				},
			},
		},
	},
}
