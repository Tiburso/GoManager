package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Tiburso/GoManager/pkg/application"
)

type Company *application.Company
type Application *application.Application

func CreateCompany(companies map[string]Company, name string) error {
	var candidatePortal string

	fmt.Print("Enter company candidate portal: ")
	fmt.Scanln(&candidatePortal)

	company, err := application.NewCompany(name, candidatePortal)

	if err != nil {
		return err
	}

	companies[name] = company

	return nil
}

func CreateApplication(companies map[string]Company, applications map[string]Application) error {

	var name, applicationType, applicationDate, companyName string

	fmt.Print("Enter application name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter application type: ")
	fmt.Scanln(&applicationType)
	fmt.Print("Enter application date: ")
	fmt.Scanln(&applicationDate)
	fmt.Print("Enter company name: ")
	fmt.Scanln(&companyName)

	_, ok := companies[companyName]

	if !ok {
		fmt.Println("Company not found!")
		CreateCompany(companies, companyName)
	}

	return nil
}

func DeleteApplication() {

}

func UpdateApplication() {

}

func ShowApplications() {

}

func ShowCompanies() {

}

func ShowMenu() {
	fmt.Println("1. Create application")
	fmt.Println("2. Delete application")
	fmt.Println("3. Update application")
	fmt.Println("4. Show applications")
	fmt.Println("5. Show companies")
	fmt.Println("6. Exit")
	fmt.Print("Enter your choice: ")
}

func main() {
	companies := make(map[string]Company)
	applications := make(map[string]Application)

	for {
		// show menu
		ShowMenu()

		// read user input from stdin
		var input string
		fmt.Scanln(&input)

		input = strings.TrimSpace(input)

		// switch on user input
		switch input {
		case "1":
			CreateApplication(companies, applications)
		case "2":
			DeleteApplication()
		case "3":
			UpdateApplication()
		case "4":
			ShowApplications()
		case "5":
			ShowCompanies()
		case "6":
			os.Exit(0)
		default:
			fmt.Println("Invalid input")
		}
	}
}
