package structs

type Application struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Status          string `json:"status"`
	ApplicationDate string `json:"application_date"`

	// Company *Company `json:"company_name,omitempty"`
}
