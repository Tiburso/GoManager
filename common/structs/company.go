package structs

type Company struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	CandidatePortal string `json:"candidate_portal"`

	Applications []*Application `json:"applications,omitempty"`
}
