package cmd

import (
	"errors"
	"log"

	"github.com/urfave/cli/v2"
)

func CreateApplication(cCtx *cli.Context) error {
	res, err := ApiRequest(&Request{
		Protocol: "POST",
		Endpoint: "/application",
		Body: map[string]any{
			"name":             cCtx.String("name"),
			"company_name":     cCtx.String("company"),
			"type":             cCtx.String("type"),
			"application_date": cCtx.String("date"),
		},
		Headers:     map[string]string{},
		QueryParams: map[string]string{},
	})

	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		return errors.New("application creation failed")
	}
	defer res.Body.Close()

	log.Println("Application created successfully")

	return nil
}

func CreateApplicationSubCommand() *cli.Command {
	return &cli.Command{
		Name:        "create",
		Aliases:     []string{"c"},
		Usage:       "create application",
		Description: "create application",
		Action:      CreateApplication,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "application name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "company",
				Aliases:  []string{"c"},
				Usage:    "company name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "type",
				Aliases:  []string{"t"},
				Usage:    "application type",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "date",
				Aliases:  []string{"d"},
				Usage:    "application date",
				Required: true,
			},
		},
	}
}

func EditApplication(cCtx *cli.Context) error {
	body := map[string]any{}

	body["name"] = cCtx.String("name")
	body["company_name"] = cCtx.String("company")

	if cCtx.IsSet("type") {
		body["type"] = cCtx.String("type")
	}

	if cCtx.IsSet("date") {
		body["application_date"] = cCtx.String("date")
	}

	if cCtx.IsSet("status") {
		body["status"] = cCtx.String("status")
	}

	res, err := ApiRequest(&Request{
		Protocol:    "PUT",
		Endpoint:    "/application",
		Body:        body,
		Headers:     map[string]string{},
		QueryParams: map[string]string{},
	})

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("application edit failed")
	}

	defer res.Body.Close()

	log.Println("Application edited successfully")

	return nil
}

func EditApplicationSubCommand() *cli.Command {
	return &cli.Command{
		Name:        "edit",
		Aliases:     []string{"e"},
		Usage:       "edit application",
		Description: "edit application",
		Action:      EditApplication,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "application name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "company",
				Aliases:  []string{"c"},
				Usage:    "company name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "type",
				Aliases:  []string{"t"},
				Usage:    "application type",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "date",
				Aliases:  []string{"d"},
				Usage:    "application date",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "status",
				Aliases:  []string{"s"},
				Usage:    "application status",
				Required: false,
			},
		},
	}
}

func DeleteApplication(cCtx *cli.Context) error {
	res, err := ApiRequest(&Request{
		Protocol: "DELETE",
		Endpoint: "/application",
		Body:     map[string]any{},
		Headers:  map[string]string{},
		QueryParams: map[string]string{
			"name": cCtx.String("name"),
		},
	})

	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		return errors.New("application deletion failed")
	}

	defer res.Body.Close()

	log.Println("Application deleted successfully")

	return nil
}

func DeleteApplicationSubCommand() *cli.Command {
	return &cli.Command{
		Name:        "delete",
		Aliases:     []string{"d"},
		Usage:       "delete application",
		Description: "delete application",
		Action:      DeleteApplication,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "application name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "company",
				Aliases:  []string{"c"},
				Usage:    "company name",
				Required: true,
			},
		},
	}
}

func GetApplications(cCtx *cli.Context) error {
	res, err := ApiRequest(&Request{
		Protocol:    "GET",
		Endpoint:    "/applications",
		Body:        map[string]any{},
		Headers:     map[string]string{},
		QueryParams: map[string]string{},
	})

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("getting applications failed")
	}

	return PrintApplications(res)
}

func GetApplicationsSubCommand() *cli.Command {
	return &cli.Command{
		Name:        "get",
		Aliases:     []string{"ga"},
		Usage:       "get applications",
		Description: "get applications",
		Action:      GetApplications,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "all",
				Aliases:  []string{"a"},
				Usage:    "get all applications",
				Required: true,
			},
		},
	}
}

func ApplicationCommand() *cli.Command {
	return &cli.Command{
		Name:        "application",
		Aliases:     []string{"a"},
		Usage:       "application commands",
		Description: "application commands",
		Subcommands: []*cli.Command{
			CreateApplicationSubCommand(),
			EditApplicationSubCommand(),
			DeleteApplicationSubCommand(),
			GetApplicationsSubCommand(),
		},
	}
}
