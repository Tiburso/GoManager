package company

import (
	"fmt"
	"regexp"

	"github.com/Tiburso/GoManager/common/structs"
	company_model "github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/models/db"
	"github.com/Tiburso/GoManager/services/convert"
)

func isValidURL(url string) bool {
	// Regular expression for a basic URL validation
	// Note: This is a simple example and might not cover all edge cases.
	urlPattern := `^(http|https)://[a-zA-Z0-9\-._~:/?#[\]@!$&'()*+,;=]+$`

	match, err := regexp.MatchString(urlPattern, url)
	if err != nil {
		// Handle error if the regular expression compilation fails
		fmt.Println("Error:", err)
		return false
	}

	return match
}

func CreateCompany(name, candidatePortal string) error {
	db := db.DB

	if name == "" {
		return fmt.Errorf("missing name")
	}

	if candidatePortal == "" || !isValidURL(candidatePortal) {
		return fmt.Errorf("missing or invalid candidate portal URL")
	}

	company := &company_model.Company{
		Name:            name,
		CandidatePortal: candidatePortal,
	}

	err := company_model.NewCompany(db, company)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCompany(name string) error {
	db := db.DB

	err := company_model.DeleteCompany(db, name)

	if err != nil {
		return err
	}

	return nil
}

func UpdateCompany(name, candidatePortal string) error {
	db := db.DB

	company, err := company_model.GetCompany(db, name)

	if err != nil {
		return err
	}

	if candidatePortal != "" {
		if !isValidURL(candidatePortal) {
			return fmt.Errorf("invalid candidate portal URL")
		}

		company.CandidatePortal = candidatePortal
	}

	_, err = company_model.UpdateCompany(db, company)

	if err != nil {
		return err
	}

	return nil
}

func GetCompanyWithApplications(name string) (*structs.CompanyWithApplications, error) {
	db := db.DB

	company, err := company_model.GetCompany(db, name)

	if err != nil {
		return nil, err
	}

	return convert.ToCompanyWithApplications(company), nil
}

func GetCompanies() ([]*structs.Company, error) {
	db := db.DB

	companies, err := company_model.GetCompanies(db)

	if err != nil {
		return nil, err
	}

	return convert.ToCompanies(companies), nil
}
