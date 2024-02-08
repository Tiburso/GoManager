package company

import (
	"fmt"
	"time"

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

func CreateApplication(name string, applicationType string, applicationDate string, companyId uint) error {
	db := db.DB

	// Application date must be a valid date
	date, err := time.Parse("2006-01-02", applicationDate)
	if err != nil {
		return fmt.Errorf("'%s' is not a valid date", applicationDate)
	}

	// Check if the application type is valid
	if !IsValidType(applicationType) {
		return company_model.ErrInvalidApplicationType{Type: applicationType}
	}

	// Check if the company exists
	company, err := company_model.GetCompany(db, companyId)

	if err != nil {
		return err
	}

	a := &company_model.Application{
		Name:            name,
		Type:            company_model.Type(applicationType),
		ApplicationDate: date,
		Status:          company_model.Applied,
		Company:         company,
	}

	return company_model.NewApplication(db, a)
}

func DeleteApplication(id uint) error {
	return company_model.DeleteApplication(db.DB, id)
}

func UpdateApplication(id uint, name, applicationType, applicationDate, applicationStatus string) error {
	db := db.DB

	application, err := company_model.GetApplication(db, id)

	if err != nil {
		return err
	}

	if applicationType != "" {
		if !IsValidType(applicationType) {
			return fmt.Errorf("'%s' is not a valid application type", applicationType)
		}

		application.Type = company_model.Type(applicationType)
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

		application.Status = company_model.Status(applicationStatus)
	}

	return company_model.UpdateApplication(db, application)
}

func GetApplication(id uint) (*company_model.Application, error) {
	return company_model.GetApplication(db.DB, id)
}

func GetApplications() ([]*company_model.Application, error) {
	return company_model.GetApplications(db.DB)
}
