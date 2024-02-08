package structs

type Company struct {
	Name            string `json:"name,omitempty"`
	CandidatePortal string `json:"candidate_portal"`

	Applications []*Application `json:"applications,omitempty"`
}
