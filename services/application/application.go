package application

import (
	"fmt"
	"time"

	application_model "github.com/Tiburso/GoManager/models/application"
	company_model "github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/models/db"
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
	_, err = company_model.GetCompany(db, companyName)

	if err != nil {
		return err
	}

	a := &application_model.Application{
		Name:            name,
		Type:            application_model.Type(applicationType),
		ApplicationDate: date,
		Status:          application_model.Applied,
		CompanyName:     companyName,
	}

	return application_model.NewApplication(db, a)
}

func DeleteApplication(name string, companyName string) error {
	return application_model.DeleteApplication(db.DB, name, companyName)
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

	return application_model.UpdateApplication(db, application)
}

func GetApplication(name string, companyName string) (*application_model.Application, error) {
	return application_model.GetApplication(db.DB, name, companyName)
}

func GetApplications() ([]*application_model.Application, error) {
	return application_model.GetApplications(db.DB)
}
