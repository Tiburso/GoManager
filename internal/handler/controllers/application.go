package controllers

/*

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Tiburso/GoManager/internal/application"
	"gorm.io/gorm"
)

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
*/
