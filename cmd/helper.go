package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tiburso/GoManager/routers/structs"
)

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
		fmt.Print(application.Company.Name + " ")
		fmt.Println(application.Company.CandidatePortal)
	}

	defer response.Body.Close()

	return nil
}

func PrintCompanies(response *http.Response) error {
	companies := []structs.Company{}

	// decode response
	err := json.NewDecoder(response.Body).Decode(&companies)

	if err != nil {
		return err
	}

	// print companies
	fmt.Println("Name CandidatePortal")
	for _, company := range companies {
		fmt.Print(company.Name + " ")
		fmt.Println(company.CandidatePortal)
	}

	defer response.Body.Close()

	return nil
}

func PrintCompany(response *http.Response) error {
	company := structs.CompanyWithApplications{}

	// decode response
	err := json.NewDecoder(response.Body).Decode(&company)

	if err != nil {
		return err
	}

	// print company
	fmt.Println("Name CandidatePortal")
	fmt.Print(company.Name + " ")
	fmt.Println(company.CandidatePortal)

	for _, application := range company.Applications {
		fmt.Print(application.Name + " ")
		fmt.Print(application.Type + " ")
		fmt.Print(application.Status + " ")
		fmt.Print(application.ApplicationDate + " ")
	}

	defer response.Body.Close()

	return nil
}
