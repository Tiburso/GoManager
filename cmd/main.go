package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Tiburso/GoManager/pkg/application"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateCompany(db *gorm.DB, name string) error {
	var candidatePortal string

	fmt.Print("Enter company candidate portal: ")
	fmt.Scanln(&candidatePortal)

	company, err := application.NewCompany(name, candidatePortal)

	if err != nil {
		return err
	}

	res := db.Create(&company)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func CreateApplication(db *gorm.DB) error {

	var name, applicationType, applicationDate, companyName string

	fmt.Print("Enter application name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter application type: ")
	fmt.Scanln(&applicationType)
	fmt.Print("Enter application date: ")
	fmt.Scanln(&applicationDate)
	fmt.Print("Enter company name: ")
	fmt.Scanln(&companyName)

	res := db.Find(&application.Company{}, "name = ?", companyName)

	if res.Error != nil {
		return res.Error
	}

	// if no rows affected, create company
	if res.RowsAffected == 0 {
		fmt.Println("Company not found!")
		err := CreateCompany(db, companyName)

		if err != nil {
			return err
		}
	}

	var company application.Company
	res = db.Find(&company, "name = ?", companyName)

	if res.Error != nil {
		return res.Error
	}

	application, err := application.NewApplication(name, applicationType, applicationDate, company)

	if err != nil {
		return err
	}

	res = db.Create(&application)

	if res.Error != nil {
		return res.Error
	}

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

func ShowApplications(db *gorm.DB) {
	var applications []application.Application
	res := db.Find(&applications)

	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}

	for _, application := range applications {
		fmt.Println(application)
	}
}

func ShowCompanies(db *gorm.DB) {
	var companies []application.Company
	res := db.Find(&companies)

	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}

	for _, company := range companies {
		fmt.Println(company)
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

func CLITool(db *gorm.DB) {
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
			CreateApplication(db)
		case "2":
			// DeleteApplication(db)
		case "3":
			// UpdateApplication(db)
		case "4":
			ShowApplications(db)
		case "5":
			ShowCompanies(db)
		case "6":
			os.Exit(0)
		default:
			fmt.Println("Invalid input")
		}
	}
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=gomanager port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&application.Company{}, &application.Application{})

	if err != nil {
		panic("failed to migrate database")
	}

	CLITool(db)
}
