package company

import (
	"fmt"

	"github.com/Tiburso/GoManager/common"
	"gorm.io/gorm"
)

type Company struct {
	*gorm.Model
	Name            string `gorm:"uniqueIndex"`
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
	id uint
}

func (e ErrCompanyNotFound) Error() string {
	return fmt.Sprintf("company with id %d does not exist", e.id)
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

func DeleteCompany(db *gorm.DB, id uint) error {
	// Delete company
	res := db.Delete(&Company{}, id)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return ErrCompanyNotFound{id: id}
	}

	return nil
}

func UpdateCompany(db *gorm.DB, c *Company) (*Company, error) {
	res := db.Save(c)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, ErrCompanyNotFound{id: c.ID}
	}

	return c, nil
}

func GetCompany(db *gorm.DB, id uint) (*Company, error) {
	var c Company
	res := db.Limit(1).Where("id = ?", id).Find(&c)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, ErrCompanyNotFound{id: id}
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
