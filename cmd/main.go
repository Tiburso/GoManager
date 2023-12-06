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

	// checking if application and company already exists

	fmt.Print("Enter application name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter company name: ")
	fmt.Scanln(&companyName)

	// first check if application already exists
	var app application.Application
	// check with name and company name
	res := db.Limit(1).Find(&app, "name = ? AND company_name = ?", name, companyName)

	if res.Error != nil {
		return res.Error
	}

	if app.Name != "" {
		return fmt.Errorf("application already exists")
	}

	res = db.Limit(1).Find(&application.Company{}, "name = ?", companyName)

	if res.Error != nil {
		return nil
	}

	if res.RowsAffected == 0 {
		err := CreateCompany(db, companyName)

		if err != nil {
			return err
		}

	} else {
		return res.Error
	}

	var company application.Company
	res = db.First(&company, "name = ?", companyName)

	if res.Error != nil {
		return res.Error
	}

	// Second Part creating the actual application
	fmt.Print("Enter application type: ")
	fmt.Scanln(&applicationType)
	fmt.Print("Enter application date: ")
	fmt.Scanln(&applicationDate)

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

func DeleteApplication(db *gorm.DB) error {
	var name, companyName string

	fmt.Print("Enter application name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter company name: ")
	fmt.Scanln(&companyName)

	res := db.Where("name = ? AND company_name = ?", name, companyName).Delete(&application.Application{})

	if res.Error != nil {
		return res.Error
	}

	return nil
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

func ShowCompanyApplications(db *gorm.DB) error {
	var companyName string

	fmt.Print("Enter company name: ")
	fmt.Scanln(&companyName)

	var company application.Company
	res := db.Limit(1).Find(&company, "name = ?", companyName)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("company does not exist")
	}

	var applications []application.Application
	res = db.Find(&applications, "company_name = ?", companyName)

	if res.Error != nil {
		return res.Error
	}

	fmt.Println(company)
	for _, application := range applications {
		fmt.Println(application)
	}

	return nil
}

func ShowMenu() {
	fmt.Println("1. Create application")
	fmt.Println("2. Delete application")
	fmt.Println("3. Update application")
	fmt.Println("4. Show applications")
	fmt.Println("5. Show companies")
	fmt.Println("6. Show company applications")
	fmt.Println("7. Exit")
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
		var err error
		switch input {
		case "1":
			err = CreateApplication(db)
		case "2":
			err = DeleteApplication(db)
		case "3":
			// err = UpdateApplication(db)
		case "4":
			ShowApplications(db)
		case "5":
			ShowCompanies(db)
		case "6":
			err = ShowCompanyApplications(db)
		case "7":
			os.Exit(0)
		default:
			fmt.Println("Invalid input")
		}

		if err != nil {
			fmt.Println(err)
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
