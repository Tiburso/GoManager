package company

import (
	"fmt"

	"gorm.io/gorm"
)

type Company struct {
	*gorm.Model
	Name            string `gorm:"primaryKey"`
	CandidatePortal string
}

// COMPANY

func NewCompany(db *gorm.DB, c *Company) error {
	// Check if company already exists
	res := db.Limit(1).Where("name = ?", "Test Company").Find(&Company{})
	if res.Error != nil {
		return res.Error
	}

	// If company already exists, return error
	if res.RowsAffected > 0 {
		return fmt.Errorf("company already exists")
	}

	// Create company
	res = db.Create(&c)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func DeleteCompany(db *gorm.DB, name string) error {
	res := db.Where("name = ?", name).Delete(&Company{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return nil
	}

	return nil
}

func UpdateCompany(db *gorm.DB, c *Company) (*Company, error) {
	res := db.Where("name = ?", c.Name).Updates(&c)

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
		return nil, nil
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
