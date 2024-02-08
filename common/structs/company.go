package structs

type Company struct {
	Name            string `json:"name"`
	CandidatePortal string `json:"candidate_portal"`

	Applications []*Application `json:"applications,omitempty"`
}
