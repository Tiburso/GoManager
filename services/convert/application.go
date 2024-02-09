package convert

import (
	"github.com/Tiburso/GoManager/common/structs"
	typeconversions "github.com/Tiburso/GoManager/common/type_conversions"
	"github.com/Tiburso/GoManager/models/company"
)

func ToApplication(a *company.Application) *structs.Application {
	return &structs.Application{
		ID:              typeconversions.ConvertToString(a.ID),
		Name:            a.Name,
		Type:            string(a.Type),
		Status:          string(a.Status),
		ApplicationDate: a.ApplicationDate.Format("2006-01-02"),
	}
}

func ToApplications(a []*company.Application) []*structs.Application {
	res := make([]*structs.Application, len(a))

	for i, app := range a {
		res[i] = ToApplication(app)
	}

	return res
}
