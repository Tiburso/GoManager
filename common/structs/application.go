package structs

type Application struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	ApplicationDate string `json:"application_date"`
	Status          string `json:"status,omitempty"`

	CompanyID uint `json:"company_id,omitempty"`
}
