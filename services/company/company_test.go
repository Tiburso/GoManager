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

	company, err := CreateCompany("Test Company", "https://www.testcompany.com/careers")

	assert.NoError(t, err)
	unittest.AssertExists(t, company)
}

func TestNewDuplicateCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	_, err := CreateCompany("Company 1", "https://www.testcompany.com/careers")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrDuplicateCompany{}, err)
	}
}

func TestNewInvalidURLCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	_, err := CreateCompany("Test Company", "invalid_url")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrCompanyInvalidURL{}, err)
	}

	_, err = CreateCompany("Test Company", "")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrCompanyInvalidURL{}, err)
	}
}

func TestDeleteCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	c, err := company_model.GetCompany(db.DB, 1)

	assert.NoError(t, err)
	unittest.AssertExists(t, c)

	err = DeleteCompany(1)

	assert.NoError(t, err)

	unittest.AssertNotExists(t, c)
}

func TestDeleteNonExistingCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := DeleteCompany(9999)

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrCompanyNotFound{}, err)
	}
}

func TestUpdateCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	c, err := company_model.GetCompany(db.DB, 1)

	assert.NoError(t, err)
	unittest.AssertExists(t, c)

	err = UpdateCompany(1, "https://www.testcompany.com/jobs")

	assert.NoError(t, err)

	c, err = company_model.GetCompany(db.DB, 1)

	assert.NoError(t, err)

	assert.Equal(t, "https://www.testcompany.com/jobs", c.CandidatePortal)
}

func TestUpdateNonExistingCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := UpdateCompany(9999, "https://www.testcompany.com/jobs")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrCompanyNotFound{}, err)
	}
}

func TestGetCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	c, err := GetCompany(1)

	assert.NoError(t, err)
	unittest.AssertExists(t, c)
}

func TestGetCompanies(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	companies, err := GetCompanies()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(companies))
}
