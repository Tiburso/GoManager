package main

import (
	"log"
	"os"

	"github.com/Tiburso/GoManager/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "GoManager",
		Usage:                "GoManager CLI",
		Commands: []*cli.Command{
			cmd.ServerCommand(),
			cmd.ClientCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
