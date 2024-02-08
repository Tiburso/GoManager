package convert

import (
	"github.com/Tiburso/GoManager/common/structs"
	company_model "github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/models/db"
)

func ToCompany(c *company_model.Company) *structs.Company {
	return &structs.Company{
		Name:            c.Name,
		CandidatePortal: c.CandidatePortal,
	}
}

func ToCompanies(c []*company_model.Company) []*structs.Company {
	res := make([]*structs.Company, len(c))

	for i, company := range c {
		res[i] = ToCompany(company)
	}

	return res
}

func ToCompanyWithApplications(c *company_model.Company) *structs.Company {
	company := ToCompany(c)

	applications, _ := company_model.GetCompanyApplications(db.DB, c.ID)

	company.Applications = ToApplications(applications)

	return company
}
