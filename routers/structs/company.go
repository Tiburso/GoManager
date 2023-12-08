package structs

type Company struct {
	Name            string `json:"name"`
	CandidatePortal string `json:"candidate_portal"`
}

type CompanyWithApplications struct {
	Company
	Applications []*ApplicationCreation `json:"applications"`
}
