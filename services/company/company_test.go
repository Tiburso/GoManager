package company

import (
	"testing"

	company_model "github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/models/db"
	"github.com/Tiburso/GoManager/models/unittest"
	"github.com/stretchr/testify/assert"
)

func TestNewCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := CreateCompany("Test Company", "https://www.testcompany.com/careers")

	assert.NoError(t, err)

	c, err := company_model.GetCompany(db.DB, "Test Company")

	assert.NoError(t, err)
	unittest.AssertExists(t, c)
}

func TestNewDuplicateCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := CreateCompany("Company 1", "https://www.testcompany.com/careers")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrDuplicateCompany{}, err)
	}
}

func TestNewInvalidURLCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := CreateCompany("Test Company", "invalid_url")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrCompanyInvalidURL{}, err)
	}

	c, err := company_model.GetCompany(db.DB, "Test Company")

	assert.Nil(t, c)
	assert.IsType(t, company_model.ErrCompanyNotFound{}, err)
}

func TestDeleteCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	c, err := company_model.GetCompany(db.DB, "Company 1")

	assert.NoError(t, err)
	unittest.AssertExists(t, c)

	err = DeleteCompany("Company 1")

	assert.NoError(t, err)

	unittest.AssertNotExists(t, c)
}

func TestDeleteNonExistingCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := DeleteCompany("Test Company")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrCompanyNotFound{}, err)
	}
}

func TestUpdateCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	c, err := company_model.GetCompany(db.DB, "Company 1")

	assert.NoError(t, err)
	unittest.AssertExists(t, c)

	err = UpdateCompany("Company 1", "https://www.testcompany.com/jobs")

	assert.NoError(t, err)

	c, err = company_model.GetCompany(db.DB, "Company 1")

	assert.NoError(t, err)

	assert.Equal(t, "https://www.testcompany.com/jobs", c.CandidatePortal)
}

func TestUpdateNonExistingCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := UpdateCompany("Test Company", "https://www.testcompany.com/jobs")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrCompanyNotFound{}, err)
	}
}

func TestGetCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	c, err := GetCompanyWithApplications("Company 1")

	assert.NoError(t, err)
	unittest.AssertExists(t, c)
}

func TestGetCompanies(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	companies, err := GetCompanies()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(companies))
}
