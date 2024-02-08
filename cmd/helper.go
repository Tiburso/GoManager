package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tiburso/GoManager/common"
	"github.com/Tiburso/GoManager/common/structs"
)

type Request struct {
	Protocol string
	Endpoint string

	Body        map[string]any
	Headers     map[string]string
	QueryParams map[string]string
}

func PrintApplications(response *http.Response) error {
	applications := []structs.Application{}

	// decode response
	err := json.NewDecoder(response.Body).Decode(&applications)

	if err != nil {
		return err
	}

	// print applications
	fmt.Println("Name Type Status ApplicationDate CompanyName CandidatePortal")
	for _, application := range applications {
		fmt.Print(application.Name + " ")
		fmt.Print(application.Type + " ")
		fmt.Print(application.Status + " ")
		fmt.Print(application.ApplicationDate + " ")
	}

	defer response.Body.Close()

	return nil
}

func printSeparator(size int) {
	for i := 0; i < size; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func PrintCompanies(response *http.Response) error {
	companies := []structs.Company{}

	// decode response
	err := json.NewDecoder(response.Body).Decode(&companies)

	if err != nil {
		return err
	}

	// get max string length for the name
	// and for the candidate portal
	maxNameLength := 0
	maxCandidatePortalLength := 0

	for _, company := range companies {
		if len(company.Name) > maxNameLength {
			maxNameLength = len(company.Name)
		}

		if len(company.CandidatePortal) > maxCandidatePortalLength {
			maxCandidatePortalLength = len(company.CandidatePortal)
		}
	}

	// print header
	fmt.Print("| Name ")
	for i := 0; i < maxNameLength-len("Name"); i++ {
		fmt.Print(" ")
	}
	fmt.Print("| CandidatePortal")
	for i := 0; i < maxCandidatePortalLength-len("CandidatePortal"); i++ {
		fmt.Print(" ")
	}

	fmt.Println(" |")

	// print separator
	printSeparator(maxNameLength + maxCandidatePortalLength + 7)

	// print companies
	for _, company := range companies {
		fmt.Print("| " + company.Name + " ")
		for i := 0; i < maxNameLength-len(company.Name); i++ {
			fmt.Print(" ")
		}
		fmt.Print("| " + company.CandidatePortal)
		for i := 0; i < maxCandidatePortalLength-len(company.CandidatePortal); i++ {
			fmt.Print(" ")
		}
		fmt.Println(" |")

		// 3 is the extra space being printed between the columns
		printSeparator(maxNameLength + maxCandidatePortalLength + 7)
	}

	defer response.Body.Close()

	return nil
}

func PrintCompany(response *http.Response) error {
	company := structs.Company{}

	// decode response
	err := json.NewDecoder(response.Body).Decode(&company)

	if err != nil {
		return err
	}

	for _, application := range company.Applications {
		fmt.Print(application.Name + " ")
		fmt.Print(application.Type + " ")
		fmt.Print(application.Status + " ")
		fmt.Println(application.ApplicationDate + " ")
	}

	defer response.Body.Close()

	return nil
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
