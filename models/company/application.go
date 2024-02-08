package company

import (
	"fmt"
	"time"

	"github.com/Tiburso/GoManager/common"
	"gorm.io/gorm"
)

type ErrApplicationNotFound struct {
	id uint
}

func (e ErrApplicationNotFound) Error() string {
	return fmt.Sprintf("application with id %d does not exist", e.id)
}

func (e ErrApplicationNotFound) Unwrap() error {
	return common.ErrNotExist
}

// Invalid arguments
type ErrInvalidApplicationType struct {
	Type string
}

func (e ErrInvalidApplicationType) Error() string {
	return fmt.Sprintf("invalid application type %s", e.Type)
}

func (e ErrInvalidApplicationType) Unwrap() error {
	return common.ErrInvalidArgument
}

type ErrInvalidApplicationStatus struct {
	Status string
}

func (e ErrInvalidApplicationStatus) Error() string {
	return fmt.Sprintf("invalid application status %s", e.Status)
}

func (e ErrInvalidApplicationStatus) Unwrap() error {
	return common.ErrInvalidArgument
}

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
	Name            string `gorm:"index"`
	Type            Type
	Status          Status
	ApplicationDate time.Time

	Company   *Company
	CompanyID uint
}

func NewApplication(db *gorm.DB, app *Application) error {
	// create application
	res := db.Create(&app)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func DeleteApplication(db *gorm.DB, id uint) error {
	res := db.Delete(&Application{}, id)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return ErrApplicationNotFound{id: id}
	}

	return nil
}

func UpdateApplication(db *gorm.DB, a *Application) error {
	res := db.Save(&a)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return ErrApplicationNotFound{id: a.ID}
	}

	return nil
}

func GetApplication(db *gorm.DB, id uint) (*Application, error) {
	var app Application

	res := db.Preload("Company").Limit(1).Where("id = ?", id).Find(&app)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, ErrApplicationNotFound{id: id}
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

func GetCompanyApplications(db *gorm.DB, company_id uint) ([]*Application, error) {
	var apps []*Application

	res := db.Where("company_id = ?", company_id).Find(&apps)

	if res.Error != nil {
		return nil, res.Error
	}

	return apps, nil
}
