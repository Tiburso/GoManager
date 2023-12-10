package convert

import (
	"github.com/Tiburso/GoManager/common/structs"
	"github.com/Tiburso/GoManager/models/application"
	"github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/models/db"
)

func ToCompany(c *company.Company) *structs.Company {
	return &structs.Company{
		Name:            c.Name,
		CandidatePortal: c.CandidatePortal,
	}
}

func ToCompanies(c []*company.Company) []*structs.Company {
	res := make([]*structs.Company, len(c))

	for i, company := range c {
		res[i] = ToCompany(company)
	}

	return res
}

func ToCompanyWithApplications(c *company.Company) *structs.CompanyWithApplications {
	applications, _ := application.GetCompanyApplications(db.DB, c.Name)

	return &structs.CompanyWithApplications{
		Company:      *ToCompany(c),
		Applications: ToApplicationCreations(applications),
	}
}
