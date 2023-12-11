package company

import (
	"testing"

	company_model "github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/models/db"
	"github.com/Tiburso/GoManager/models/unittest"
	"github.com/stretchr/testify/assert"
)

func TestNewValidCompany(t *testing.T) {
	err := CreateCompany("Test Company", "https://www.testcompany.com/careers")

	assert.NoError(t, err)

	c, err := company_model.GetCompany(db.DB, "Test Company")

	assert.NoError(t, err)
	unittest.AssertExists(t, c)
}

func TestNewDuplicateCompany(t *testing.T) {
	err := CreateCompany("Test Company", "https://www.testcompany.com/careers")

	assert.NoError(t, err)

	c, err := company_model.GetCompany(db.DB, "Test Company")

	assert.NoError(t, err)
	unittest.AssertExists(t, c)

	err = CreateCompany("Test Company", "https://www.testcompany.com/careers")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrDuplicateCompany{}, err)
	}

	assert.Nil(t, c)
}

func TestNewInvalidURLCompany(t *testing.T) {
	err := CreateCompany("Test Company", "invalid_url")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrCompanyInvalidURL{}, err)
	}

	c, err := company_model.GetCompany(db.DB, "Test Company")

	assert.Nil(t, c)
	assert.IsType(t, company_model.ErrCompanyNotFound{}, err)
}

func TestDeleteCompany(t *testing.T) {
	err := CreateCompany("Test Company", "https://www.testcompany.com/careers")

	assert.NoError(t, err)

	c, err := company_model.GetCompany(db.DB, "Test Company")

	assert.NoError(t, err)
	unittest.AssertExists(t, c)

	err = DeleteCompany("Test Company")

	assert.NoError(t, err)

	unittest.AssertNotExists(t, c)
}

func TestDeleteNonExistingCompany(t *testing.T) {
	err := DeleteCompany("Test Company")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrCompanyNotFound{}, err)
	}
}

func TestUpdateCompany(t *testing.T) {
	err := CreateCompany("Test Company", "https://www.testcompany.com/careers")

	assert.NoError(t, err)

	c, err := company_model.GetCompany(db.DB, "Test Company")

	assert.NoError(t, err)
	unittest.AssertExists(t, c)

	err = UpdateCompany("Test Company", "https://www.testcompany.com/jobs")

	assert.NoError(t, err)

	c, err = company_model.GetCompany(db.DB, "Test Company")

	assert.NoError(t, err)

	assert.Equal(t, "https://www.testcompany.com/jobs", c.CandidatePortal)
}

func TestUpdateNonExistingCompany(t *testing.T) {
	err := UpdateCompany("Test Company", "https://www.testcompany.com/jobs")

	if assert.Error(t, err) {
		assert.IsType(t, company_model.ErrCompanyNotFound{}, err)
	}
}
