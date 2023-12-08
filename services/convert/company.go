package convert

import (
	"github.com/Tiburso/GoManager/models/application"
	"github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/routers/structs"
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

func ToCompanyWithApplications(c *company.Company, applications []*application.Application) *structs.CompanyWithApplications {
	return &structs.CompanyWithApplications{
		Company:      *ToCompany(c),
		Applications: ToApplicationCreations(applications),
	}
}
