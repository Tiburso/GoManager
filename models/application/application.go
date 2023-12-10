package application

import (
	"fmt"
	"time"

	"github.com/Tiburso/GoManager/models/company"
	"gorm.io/gorm"
)

type Type string

const (
	FullTime   Type = "full_time"
	PartTime   Type = "part_time"
	Internship Type = "internship"
)

type Status string

const (
	Applied  Status = "applied"
	Rejected Status = "rejected"
	Accepted Status = "accepted"
)

type Application struct {
	*gorm.Model
	Name            string `gorm:"primaryKey"`
	Type            Type
	Status          Status
	ApplicationDate time.Time

	CompanyName string `gorm:"primaryKey"`
	Company     *company.Company
}

func NewApplication(db *gorm.DB, app *Application) error {
	// check if application already exists
	res := db.Limit(1).Where("name = ? AND company_name = ?", app.Name, app.CompanyName).Find(&Application{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected > 0 {
		return fmt.Errorf("application already exists")
	}

	// create application
	res = db.Create(&app)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func DeleteApplication(db *gorm.DB, name string, companyName string) error {
	res := db.Where("name = ? AND company_name = ?", name, companyName).Delete(&Application{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("application does not exist")
	}

	return nil
}

func UpdateApplication(db *gorm.DB, a *Application) error {
	res := db.Where("name = ? AND company_name = ?", a.Name, a.CompanyName).Updates(&a)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("application does not exist")
	}

	return nil
}

func GetApplication(db *gorm.DB, name string, companyName string) (*Application, error) {
	var app Application

	res := db.Limit(1).Find(&app, "name = ? AND company_name = ?", name, companyName)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("application does not exist")
	}

	return &app, nil
}

func GetApplications(db *gorm.DB) ([]*Application, error) {
	var apps []*Application

	res := db.Model(&Application{}).Preload("Company").Find(&apps)

	if res.Error != nil {
		return nil, res.Error
	}

	return apps, nil
}

func GetCompanyApplications(db *gorm.DB, companyName string) ([]*Application, error) {
	var apps []*Application

	res := db.Find(&apps, "company_name = ?", companyName)

	if res.Error != nil {
		return nil, res.Error
	}

	return apps, nil
}
