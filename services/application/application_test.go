package application

import (
	"testing"

	application_model "github.com/Tiburso/GoManager/models/application"
	"github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/models/db"
	"github.com/Tiburso/GoManager/models/unittest"
	"github.com/stretchr/testify/assert"
)

func TestCreateApplication(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := CreateApplication(
		"Test Company",
		string(application_model.Internship),
		"2023-01-01",
		"Company 1")

	assert.NoError(t, err)

	a, err := application_model.GetApplication(db.DB, "Test Company", "Company 1")

	assert.NoError(t, err)
	unittest.AssertExists(t, a)
}

func TestCreateDuplicateApplication(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := CreateApplication("Application 1",
		string(application_model.Internship),
		"2023-01-01",
		"Company 1")

	if assert.Error(t, err) {
		assert.IsType(t, application_model.ErrDuplicateApplication{}, err)
	}
}

func TestCreateApplicationWithInvalidType(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := CreateApplication("Test Company",
		"invalid type",
		"2023-01-01",
		"Company 1")

	if assert.Error(t, err) {
		assert.IsType(t, application_model.ErrInvalidApplicationType{}, err)
	}
}

func TestCreateApplicationWithInvalidStatus(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := CreateApplication("Test Company",
		string(application_model.Internship),
		"2023-01-01",
		"invalid status")

	if assert.Error(t, err) {
		assert.IsType(t, application_model.ErrInvalidApplicationStatus{}, err)
	}
}

func TestCreateApplicationWithNoCompany(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := CreateApplication("Test Company",
		string(application_model.Internship),
		"2023-01-01",
		"")

	if assert.Error(t, err) {
		assert.IsType(t, company.ErrCompanyNotFound{}, err)
	}
}

func TestDeleteApplication(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := DeleteApplication("Application 1", "Company 1")

	assert.NoError(t, err)

	a, err := application_model.GetApplication(db.DB, "Application 1", "Company 1")

	assert.NoError(t, err)
	unittest.AssertNotExists(t, a)
}

func TestDeleteApplicationInvalid(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := DeleteApplication("Application 1", "")

	if assert.Error(t, err) {
		assert.IsType(t, application_model.ErrApplicationNotFound{}, err)
	}

	err = DeleteApplication("", "Company 1")

	if assert.Error(t, err) {
		assert.IsType(t, application_model.ErrApplicationNotFound{}, err)
	}
}

func TestUpdateApplication(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := UpdateApplication("Application 1",
		string(application_model.Internship),
		"2023-01-01",
		string(application_model.Applied),
		"Company 1")

	assert.NoError(t, err)

	a, err := application_model.GetApplication(db.DB, "Application 1", "Company 1")

	assert.NoError(t, err)
	unittest.AssertExists(t, a)
}

func TestUpdateApplicationInvalid(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	err := UpdateApplication("Application 1",
		string(application_model.Internship),
		"2023-01-01",
		string(application_model.Applied),
		"")

	if assert.Error(t, err) {
		assert.IsType(t, company.ErrCompanyNotFound{}, err)
	}

	err = UpdateApplication("Application 1",
		string(application_model.Internship),
		"2023-01-01",
		string(application_model.Applied),
		"Company 2")

	if assert.Error(t, err) {
		assert.IsType(t, application_model.ErrApplicationNotFound{}, err)
	}
}

func TestGetApplication(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	a, err := GetApplication("Application 1", "Company 1")

	assert.NoError(t, err)
	assert.NotNil(t, a)
}

func TestGetApplicationInvalid(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	a, err := GetApplication("Application 1", "")

	if assert.Error(t, err) {
		assert.IsType(t, application_model.ErrApplicationNotFound{}, err)
	}

	assert.Nil(t, a)

	a, err = GetApplication("", "Company 1")

	if assert.Error(t, err) {
		assert.IsType(t, application_model.ErrApplicationNotFound{}, err)
	}

	assert.Nil(t, a)
}

func TestGetApplications(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())

	a, err := GetApplications()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(a))
}
