package application

import (
	"fmt"
	"time"

	"github.com/Tiburso/GoManager/common/structs"
	application_model "github.com/Tiburso/GoManager/models/application"
	company_model "github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/models/db"
	"github.com/Tiburso/GoManager/services/convert"
)

func IsValidType(t string) bool {
	switch t {
	case "full_time", "part_time", "internship":
		return true
	}

	return false
}

func IsValidStatus(s string) bool {
	switch s {
	case "applied", "rejected", "accepted":
		return true
	}

	return false
}

func CreateApplication(name string, applicationType string, applicationDate string, companyName string) error {
	db := db.DB

	// Application date must be a valid date
	date, err := time.Parse("2006-01-02", applicationDate)
	if err != nil {
		return fmt.Errorf("'%s' is not a valid date", applicationDate)
	}

	// Check if the application type is valid
	if !IsValidType(applicationType) {
		return application_model.ErrInvalidApplicationType{Type: applicationType}
	}

	// Check if the company exists
	company, err := company_model.GetCompany(db, companyName)

	if err != nil {
		return err
	}

	a := &application_model.Application{
		Name:            name,
		Type:            application_model.Type(applicationType),
		ApplicationDate: date,
		Status:          application_model.Applied,
		Company:         company,
	}

	err = application_model.NewApplication(db, a)

	if err != nil {
		return err
	}

	return nil
}

func DeleteApplication(name string, companyName string) error {
	db := db.DB

	err := application_model.DeleteApplication(db, name, companyName)

	if err != nil {
		return err
	}

	return nil
}

func UpdateApplication(name, applicationType, applicationDate, applicationStatus, companyName string) error {
	db := db.DB

	application, err := application_model.GetApplication(db, name, companyName)

	if err != nil {
		return err
	}

	if applicationType != "" {
		if !IsValidType(applicationType) {
			return fmt.Errorf("'%s' is not a valid application type", applicationType)
		}

		application.Type = application_model.Type(applicationType)
	}

	if applicationDate != "" {
		date, err := time.Parse("2006-01-02", applicationDate)
		if err != nil {
			return fmt.Errorf("'%s' is not a valid date", applicationDate)
		}
		application.ApplicationDate = date
	}

	if applicationStatus != "" {
		if !IsValidStatus(applicationStatus) {
			return fmt.Errorf("'%s' is not a valid application status", applicationStatus)
		}

		application.Status = application_model.Status(applicationStatus)
	}

	err = application_model.UpdateApplication(db, application)

	if err != nil {
		return err
	}

	return nil
}

func GetApplication(name string, companyName string) (*structs.Application, error) {
	db := db.DB

	app, err := application_model.GetApplication(db, name, companyName)

	if err != nil {
		return nil, err
	}

	return convert.ToApplication(app), nil
}

func GetApplications() ([]*structs.Application, error) {
	db := db.DB

	apps, err := application_model.GetApplications(db)

	if err != nil {
		return nil, err
	}

	return convert.ToApplications(apps), nil
}
