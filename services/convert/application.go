package convert

import (
	"github.com/Tiburso/GoManager/common/structs"
	"github.com/Tiburso/GoManager/models/company"
)

func ToApplicationCreation(a *company.Application) *structs.ApplicationCreation {
	return &structs.ApplicationCreation{
		Name:            a.Name,
		CompanyName:     a.Company.Name,
		Type:            string(a.Type),
		Status:          string(a.Status),
		ApplicationDate: a.ApplicationDate.Format("2006-01-02"),
	}
}

func ToApplicationCreations(a []*company.Application) []*structs.ApplicationCreation {
	res := make([]*structs.ApplicationCreation, len(a))

	for i, app := range a {
		res[i] = ToApplicationCreation(app)
	}

	return res
}

func ToApplication(a *company.Application) *structs.Application {
	return &structs.Application{
		ApplicationCreation: ToApplicationCreation(a),
		Company:             ToCompany(a.Company),
	}
}

func ToApplications(a []*company.Application) []*structs.Application {
	res := make([]*structs.Application, len(a))

	for i, app := range a {
		res[i] = ToApplication(app)
	}

	return res
}
