package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Tiburso/GoManager/pkg/application"
)

func CreateCompany(companies map[string]*application.Company, name string) error {
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

func CreateApplication(companies map[string]*application.Company, applications map[string]*application.Application) error {

	var name, applicationType, applicationDate, companyName string

	fmt.Print("Enter application name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter application type: ")
	fmt.Scanln(&applicationType)
	fmt.Print("Enter application date: ")
	fmt.Scanln(&applicationDate)
	fmt.Print("Enter company name: ")
	fmt.Scanln(&companyName)

	if _, ok := companies[companyName]; !ok {
		fmt.Println("Company not found!")
		err := CreateCompany(companies, companyName)

		if err != nil {
			return err
		}
	}

	company := companies[companyName]

	application, err := application.NewApplication(name, applicationType, applicationDate, company)

	if err != nil {
		return err
	}

	applications[name] = application

	return nil
}

func DeleteApplication(applications map[string]*application.Application) {
	var name string

	fmt.Print("Enter application name: ")
	fmt.Scanln(&name)

	delete(applications, name)
}

func UpdateApplication(applications map[string]*application.Application) {
	var name, updateType string

	fmt.Print("Enter application name: ")
	fmt.Scanln(&name)

	app, ok := applications[name]

	if !ok {
		fmt.Println("Application not found!")
		return
	}

	fmt.Println("1. Update name")
	fmt.Println("2. Update type")
	fmt.Println("3. Update date")
	fmt.Print("Enter what you want to update: ")
	fmt.Scanln(&updateType)

	updateType = strings.TrimSpace(updateType)

	switch updateType {
	case "1":
		var newName string
		fmt.Print("Enter new name: ")
		fmt.Scanln(&newName)
		err := app.SetName(newName)

		if err != nil {
			fmt.Println(err)
		}
	case "2":
		var newType string
		fmt.Print("Enter new type: ")
		fmt.Scanln(&newType)
		err := app.SetType(newType)

		if err != nil {
			fmt.Println(err)
		}

	case "3":
		var newDate string
		fmt.Print("Enter new date: ")
		fmt.Scanln(&newDate)
		err := app.SetApplicationDate(newDate)

		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Invalid input")
	}
}

func ShowApplications(applications map[string]*application.Application) {
	for _, application := range applications {
		fmt.Println(*application)
	}
}

func ShowCompanies(companies map[string]*application.Company) {
	for _, company := range companies {
		fmt.Println(*company)
	}
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
	companies := make(map[string]*application.Company)
	applications := make(map[string]*application.Application)

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
			DeleteApplication(applications)
		case "3":
			UpdateApplication(applications)
		case "4":
			ShowApplications(applications)
		case "5":
			ShowCompanies(companies)
		case "6":
			os.Exit(0)
		default:
			fmt.Println("Invalid input")
		}
	}
}
