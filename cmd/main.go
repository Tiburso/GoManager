package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Tiburso/GoManager/pkg/application"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ReadLine(scanner *bufio.Scanner) string {
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}

	return ""
}

func CreateCompany(db *gorm.DB, name string) error {
	var candidatePortal string

	fmt.Print("Enter company candidate portal: ")
	candidatePortal = ReadLine(bufio.NewScanner(os.Stdin))

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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter application name: ")
	name = ReadLine(scanner)
	fmt.Print("Enter company name: ")
	companyName = ReadLine(scanner)

	res := db.Limit(1).Find(&application.Application{}, "name = ? AND company_name = ?", name, companyName)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected != 0 {
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
	}

	var company application.Company
	res = db.First(&company, "name = ?", companyName)

	if res.Error != nil {
		return res.Error
	}

	// Second Part creating the actual application
	fmt.Print("Enter application type: ")
	applicationType = ReadLine(scanner)
	fmt.Print("Enter application date: ")
	applicationDate = ReadLine(scanner)

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

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter application name: ")
	name = ReadLine(scanner)
	fmt.Print("Enter company name: ")
	companyName = ReadLine(scanner)

	res := db.Where("name = ? AND company_name = ?", name, companyName).Delete(&application.Application{})

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func UpdateApplication(db *gorm.DB) error {
	var name, companyName, updateType string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter application name: ")
	name = ReadLine(scanner)
	fmt.Print("Enter company name: ")
	companyName = ReadLine(scanner)

	var app application.Application
	res := db.Limit(1).Find(&app, "name = ? AND company_name = ?", name, companyName)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("application does not exist")
	}

	fmt.Println("1. Update type")
	fmt.Println("2. Update date")
	fmt.Println("3. Update status")
	fmt.Print("Enter what you want to update: ")
	updateType = ReadLine(scanner)

	updateType = strings.TrimSpace(updateType)

	switch updateType {
	case "1":
		fmt.Print("Enter new type: ")
		newType := ReadLine(scanner)
		err := app.SetType(newType)

		if err != nil {
			return err
		}

	case "2":
		fmt.Print("Enter new date: ")
		newDate := ReadLine(scanner)
		err := app.SetApplicationDate(newDate)

		if err != nil {
			return err
		}

	case "3":
		fmt.Print("Enter new status: ")
		newStatus := ReadLine(scanner)
		err := app.SetStatus(newStatus)

		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("invalid input")
	}

	res = db.Save(&app)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func ShowApplications(db *gorm.DB) {
	var applications []application.Application
	res := db.Model(&application.Application{}).Preload("Company").Find(&applications)

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
	companyName = ReadLine(bufio.NewScanner(os.Stdin))

	var company application.Company
	res := db.Model(&application.Company{}).Preload("Applications").Limit(1).Find(&company, "name = ?", companyName)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("company does not exist")
	}

	fmt.Println(company)

	for _, application := range company.GetApplications() {
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

func CLITool(db *gorm.DB) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// show menu
		ShowMenu()

		// read user input from stdin
		input := ReadLine(scanner)

		// switch on user input
		var err error
		switch input {
		case "1":
			err = CreateApplication(db)
		case "2":
			err = DeleteApplication(db)
		case "3":
			err = UpdateApplication(db)
		case "4":
			ShowApplications(db)
		case "5":
			ShowCompanies(db)
		case "6":
			err = ShowCompanyApplications(db)
		case "7":
			return nil
		default:
			fmt.Println("Invalid input")
		}

		if err != nil {
			return err
		}
	}
}

func SetupDB() *gorm.DB {
	host := GetEnvWithDefault("DB_HOST", "localhost")
	user := GetEnvWithDefault("DB_USER", "postgres")
	password := GetEnvWithDefault("DB_PASSWORD", "postgres")
	dbname := GetEnvWithDefault("DB_NAME", "gomanager")
	port := GetEnvWithDefault("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&application.Company{}, &application.Application{})

	if err != nil {
		panic("failed to migrate database")
	}

	return db
}

func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		// write a warning to the log
		log.Println("Error loading .env file")
	}
}

func main() {
	LoadEnv()

	db := SetupDB()

	err := CLITool(db)

	if err != nil {
		panic(err)
	}
}
