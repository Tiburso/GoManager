package company

import (
	"testing"

	"github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/models/db"
	"github.com/Tiburso/GoManager/models/unittest"
	"github.com/stretchr/testify/assert"
)

func TestCreateApplication(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	a, err := CreateApplication(
		"Application 2",
		string(company.Internship),
		"2023-01-01",
		1)

	assert.NoError(t, err)
	assert.NotNil(t, a)
	unittest.AssertExists(t, a)
}

func TestCreateApplicationWithInvalidType(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	_, err := CreateApplication("Test Company",
		"invalid type",
		"2023-01-01",
		1)

	if assert.Error(t, err) {
		assert.IsType(t, company.ErrInvalidApplicationType{}, err)
	}
}

func TestCreateApplicationWithNoCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	_, err := CreateApplication("Test Company",
		string(company.Internship),
		"2023-01-01",
		99999)

	if assert.Error(t, err) {
		assert.IsType(t, company.ErrCompanyNotFound{}, err)
	}
}

func TestDeleteApplication(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := DeleteApplication(1)

	assert.NoError(t, err)

	a, err := company.GetApplication(db.DB, 1)

	if assert.Error(t, err) {
		assert.IsType(t, company.ErrApplicationNotFound{}, err)
	}
	assert.Nil(t, a)
}

func TestDeleteApplicationInvalid(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := DeleteApplication(99999)

	if assert.Error(t, err) {
		assert.IsType(t, company.ErrApplicationNotFound{}, err)
	}
}

func TestUpdateApplication(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	// Get initial version of application

	a, _ := company.GetApplication(db.DB, 1)

	// Confirm the application has status != "Accepted"
	assert.NotEqual(t, company.Accepted, a.Status)

	err := UpdateApplication(1,
		"Application 1",
		string(company.Internship),
		"2023-01-01",
		string(company.Accepted))

	assert.NoError(t, err)

	a, err = company.GetApplication(db.DB, 1)

	assert.NoError(t, err)
	unittest.AssertExists(t, a)
	assert.Equal(t, company.Accepted, a.Status)
	assert.Equal(t, "2023-01-01", a.ApplicationDate.Format("2006-01-02"))
}

func TestUpdateApplicationInvalid(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := UpdateApplication(9999,
		"Application 1",
		string(company.Internship),
		"2023-01-01",
		string(company.Applied))

	if assert.Error(t, err) {
		assert.IsType(t, company.ErrApplicationNotFound{}, err)
	}
}

func TestGetApplication(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	a, err := GetApplication(1)

	assert.NoError(t, err)
	assert.NotNil(t, a)
}

func TestGetApplicationInvalid(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	a, err := GetApplication(99999)

	if assert.Error(t, err) {
		assert.IsType(t, company.ErrApplicationNotFound{}, err)
	}

	assert.Nil(t, a)
}

func TestGetApplications(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	a, err := GetApplications()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(a))
}
