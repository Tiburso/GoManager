package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Tiburso/GoManager/common"
	"github.com/urfave/cli/v2"
)

type Request struct {
	Protocol string
	Endpoint string

	Body        map[string]any
	Headers     map[string]string
	QueryParams map[string]string
}

func GetServerUrl() string {
	protocol := common.GetEnvWithDefault("PROTOCOL", "http")
	endpoint := common.GetEnvWithDefault("ENDPOINT", "localhost")
	port := common.GetEnvWithDefault("PORT", "8080")

	return protocol + "://" + endpoint + ":" + port + "/api/v1"
}

func ApiRequest(request *Request) (*http.Response, error) {
	url := GetServerUrl() + request.Endpoint

	client := &http.Client{}

	// marshall body
	body, err := json.Marshal(request.Body)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(request.Protocol, url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	// add json headers
	req.Header.Set("Content-Type", "application/json")
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	q := req.URL.Query()

	for key, value := range request.QueryParams {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

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

	if res.StatusCode != 200 {
		return errors.New("application creation failed")
	}
	defer res.Body.Close()

	log.Println("Application created successfully")

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

	if res.StatusCode != 200 {
		return errors.New("application deletion failed")
	}

	defer res.Body.Close()

	log.Println("Application deleted successfully")

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
		Name:        "get-applications",
		Aliases:     []string{"ga"},
		Usage:       "get applications",
		Description: "get applications",
		Action:      GetApplications,
	}
}

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

	if res.StatusCode != 200 {
		return errors.New("company creation failed")
	}
	defer res.Body.Close()

	log.Println("Company created successfully")

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
