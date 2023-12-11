package application

import (
	"testing"

	"github.com/Tiburso/GoManager/models/db"
	"github.com/Tiburso/GoManager/models/unittest"

	c "github.com/Tiburso/GoManager/models/company"
)

func TestMain(m *testing.M) {
	unittest.MainTest(m)

	// Setup
	company := &c.Company{
		Name:            "Test Company",
		CandidatePortal: "https://www.testcompany.com/careers",
	}

	if err := c.NewCompany(db.DB, company); err != nil {
		panic(err)
	}

	// Run tests
	m.Run()

	// Teardown
	if err := c.DeleteCompany(db.DB, company.Name); err != nil {
		panic(err)
	}
}
