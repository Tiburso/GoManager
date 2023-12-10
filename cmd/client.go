package main

import "github.com/urfave/cli/v2"

func CreateApplication(cCtx *cli.Context) error {
	return nil
}

func CreateApplicationSubCommand() *cli.Command {
	return &cli.Command{
		Name:        "create-application",
		Aliases:     []string{"ca"},
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
	return nil
}

func EditApplicationSubCommand() *cli.Command {
	return &cli.Command{
		Name:        "edit-application",
		Aliases:     []string{"ea"},
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
	return nil
}

func DeleteApplicationSubCommand() *cli.Command {
	return &cli.Command{
		Name:        "delete-application",
		Aliases:     []string{"da"},
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
	return nil
}

func GetApplicationsSubCommand() *cli.Command {
	return &cli.Command{
		Name:        "get-applications",
		Aliases:     []string{"ga"},
		Usage:       "get applications",
		Description: "get applications",
		Action:      GetApplications,
	}
}

func CreateCompany(cCtx *cli.Context) error {
	return nil
}

func CreateCompanySubCommand() *cli.Command {
	return &cli.Command{
		Name:        "create-company",
		Aliases:     []string{"cc"},
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
}

func DeleteCompany(cCtx *cli.Context) error {
	return nil
}

func DeleteCompanySubCommand() *cli.Command {
	return &cli.Command{
		Name:        "delete-company",
		Aliases:     []string{"dc"},
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
}

func GetCompanies(cCtx *cli.Context) error {
	return nil
}

func GetCompaniesSubCommand() *cli.Command {
	return &cli.Command{
		Name:        "get-companies",
		Aliases:     []string{"gc"},
		Usage:       "get companies",
		Description: "get companies",
		Action:      GetCompanies,
	}
}

func GetCompany(cCtx *cli.Context) error {
	return nil
}

func GetCompanySubCommand() *cli.Command {
	return &cli.Command{
		Name:        "get-company",
		Aliases:     []string{"gco"},
		Usage:       "get company",
		Description: "get company",
		Action:      GetCompany,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "company name",
				Required: true,
			},
		},
	}
}

func ClientCommand() *cli.Command {
	return &cli.Command{
		Name:        "client",
		Aliases:     []string{"c"},
		Usage:       "client commands",
		Description: "client commands",
		Subcommands: []*cli.Command{
			CreateApplicationSubCommand(),
			EditApplicationSubCommand(),
			DeleteApplicationSubCommand(),
			GetApplicationsSubCommand(),
			CreateCompanySubCommand(),
			DeleteCompanySubCommand(),
			GetCompaniesSubCommand(),
			GetCompanySubCommand(),
		},
	}
}
