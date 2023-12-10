package cmd

import (
	"errors"
	"log"

	"github.com/urfave/cli/v2"
)

func CreateCompany(cCtx *cli.Context) error {
	res, err := ApiRequest(&Request{
		Protocol: "POST",
		Endpoint: "/company",
		Body: map[string]any{
			"name":             cCtx.String("name"),
			"candidate_portal": cCtx.String("portal"),
		},
		Headers:     map[string]string{},
		QueryParams: map[string]string{},
	})

	if err != nil {
		return err
	}

	if res.StatusCode != 201 {
		return errors.New("company creation failed")
	}
	defer res.Body.Close()

	log.Println("Company created successfully")

	return nil
}

var createCompanySubCommand *cli.Command = &cli.Command{
	Name:        "create",
	Aliases:     []string{"c"},
	Usage:       "create company",
	Description: "create company",
	Action:      CreateCompany,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "company name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "portal",
			Aliases:  []string{"p"},
			Usage:    "company portal",
			Required: true,
		},
	},
}

func DeleteCompany(cCtx *cli.Context) error {
	res, err := ApiRequest(&Request{
		Protocol: "DELETE",
		Endpoint: "/company",
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
		return errors.New("company deletion failed")
	}
	defer res.Body.Close()

	log.Println("Company deleted successfully")

	return nil
}

var deleteCompanySubCommand *cli.Command = &cli.Command{
	Name:        "delete",
	Aliases:     []string{"d"},
	Usage:       "delete company",
	Description: "delete company",
	Action:      DeleteCompany,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "company name",
			Required: true,
		},
	},
}

func GetCompany(cCtx *cli.Context) error {
	if cCtx.Bool("all") {
		return GetCompanies(cCtx)
	}

	if !cCtx.IsSet("name") {
		return errors.New("name is required")
	}

	res, err := ApiRequest(&Request{
		Protocol: "GET",
		Endpoint: "/company",
		Body:     map[string]any{},
		Headers:  map[string]string{},
		QueryParams: map[string]string{
			"name": cCtx.String("name"),
		},
	})

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("getting company failed")
	}

	return PrintCompany(res)
}

func GetCompanies(cCtx *cli.Context) error {
	res, err := ApiRequest(&Request{
		Protocol:    "GET",
		Endpoint:    "/companies",
		Body:        map[string]any{},
		Headers:     map[string]string{},
		QueryParams: map[string]string{},
	})

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("getting companies failed")
	}

	return PrintCompanies(res)
}

var getCompanySubCommand *cli.Command = &cli.Command{
	Name:        "get",
	Aliases:     []string{"g"},
	Usage:       "get company",
	Description: "get company",
	Action:      GetCompany,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "company name",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "all",
			Aliases:  []string{"a"},
			Usage:    "get all companies",
			Required: false,
		},
	},
}

func EditCompany(cCtx *cli.Context) error {
	res, err := ApiRequest(&Request{
		Protocol: "PUT",
		Endpoint: "/company",
		Body: map[string]any{
			"name":             cCtx.String("name"),
			"candidate_portal": cCtx.String("portal"),
		},
		Headers:     map[string]string{},
		QueryParams: map[string]string{},
	})

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("company edit failed")
	}

	defer res.Body.Close()

	log.Println("Company edited successfully")

	return nil
}

var editCompanySubCommand *cli.Command = &cli.Command{
	Name:        "edit",
	Aliases:     []string{"e"},
	Usage:       "edit company",
	Description: "edit company",
	Action:      EditCompany,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "company name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "portal",
			Aliases:  []string{"p"},
			Usage:    "company portal",
			Required: true,
		},
	},
}

var CompanyCommand *cli.Command = &cli.Command{
	Name:        "company",
	Aliases:     []string{"c"},
	Usage:       "company",
	Description: "company",
	Subcommands: []*cli.Command{
		createCompanySubCommand,
		deleteCompanySubCommand,
		getCompanySubCommand,
		editCompanySubCommand,
	},
}
