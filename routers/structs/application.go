package structs

type Application struct {
	Name            string `json:"name"`
	CompanyName     string `json:"company_name"`
	Type            string `json:"type"`
	Status          string `json:"status"`
	ApplicationDate string `json:"application_date"`
}
