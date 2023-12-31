package company

import (
	"fmt"

	"github.com/Tiburso/GoManager/common"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name            string `gorm:"primaryKey;uniqueIndex;not null"`
	CandidatePortal string
}

type ErrDuplicateCompany struct {
	Name string
}

func (e ErrDuplicateCompany) Error() string {
	return fmt.Sprintf("company with name %s already exists", e.Name)
}

func (e ErrDuplicateCompany) Unwrap() error {
	return common.ErrAlreadyExist
}

type ErrCompanyNotFound struct {
	Name string
}

func (e ErrCompanyNotFound) Error() string {
	return fmt.Sprintf("company with name %s does not exist", e.Name)
}

func (e ErrCompanyNotFound) Unwrap() error {
	return common.ErrNotExist
}

type ErrCompanyInvalidURL struct {
	URL string
}

func (e ErrCompanyInvalidURL) Error() string {
	return fmt.Sprintf("invalid URL %s", e.URL)
}

func (e ErrCompanyInvalidURL) Unwrap() error {
	return common.ErrInvalidArgument
}

// COMPANY

func NewCompany(db *gorm.DB, c *Company) error {
	// Check if company already exists
	res := db.Limit(1).Where("name = ?", c.Name).Find(&Company{})
	if res.Error != nil {
		return res.Error
	}

	// If company already exists, return error
	if res.RowsAffected > 0 {
		return ErrDuplicateCompany{Name: c.Name}
	}

	// Create company
	res = db.Create(&c)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func DeleteCompany(db *gorm.DB, name string) error {
	// Check if company already exists
	res := db.Limit(1).Where("name = ?", name).Find(&Company{})

	if res.Error != nil {
		return res.Error
	}

	// If company already exists, return error
	if res.RowsAffected == 0 {
		return ErrCompanyNotFound{Name: name}
	}

	res = db.Where("name = ?", name).Delete(&Company{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return nil
	}

	return nil
}

func UpdateCompany(db *gorm.DB, c *Company) (*Company, error) {
	// Check if company already exists
	res := db.Limit(1).Where("name = ?", c.Name).Find(&Company{})

	if res.Error != nil {
		return nil, res.Error
	}

	// If company already exists, return error
	if res.RowsAffected == 0 {
		return nil, ErrCompanyNotFound{Name: c.Name}
	}

	res = db.Where("name = ?", c.Name).Updates(&c)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, nil
	}

	return c, nil
}

func GetCompany(db *gorm.DB, name string) (*Company, error) {
	var c Company
	res := db.Where("name = ?", name).Find(&c)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, ErrCompanyNotFound{Name: name}
	}

	return &c, nil
}

func GetCompanies(db *gorm.DB) ([]*Company, error) {
	var companies []*Company
	res := db.Find(&companies)

	if res.Error != nil {
		return nil, res.Error
	}

	return companies, nil
}
