package convert

import (
	"github.com/Tiburso/GoManager/models/application"
	"github.com/Tiburso/GoManager/routers/structs"
)

func ToApplicationCreation(a *application.Application) *structs.ApplicationCreation {
	return &structs.ApplicationCreation{
		Name:            a.Name,
		CompanyName:     a.CompanyName,
		Type:            string(a.Type),
		Status:          string(a.Status),
		ApplicationDate: a.ApplicationDate.Format("2006-01-02"),
	}
}

func ToApplicationCreations(a []*application.Application) []*structs.ApplicationCreation {
	res := make([]*structs.ApplicationCreation, len(a))

	for i, app := range a {
		res[i] = ToApplicationCreation(app)
	}

	return res
}

func ToApplication(a *application.Application) *structs.Application {
	return &structs.Application{
		ApplicationCreation: ToApplicationCreation(a),
		Company:             ToCompany(a.Company),
	}
}

func ToApplications(a []*application.Application) []*structs.Application {
	res := make([]*structs.Application, len(a))

	for i, app := range a {
		res[i] = ToApplication(app)
	}

	return res
}
