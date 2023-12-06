package application

import (
	"fmt"
	"regexp"
)

type Company struct {
	Name            string `gorm:"primaryKey"`
	CandidatePortal string
	Applications    []Application `gorm:"foreignKey:CompanyName;references:Name"`
}

func isValidURL(url string) bool {
	// Regular expression for a basic URL validation
	// Note: This is a simple example and might not cover all edge cases.
	urlPattern := `^(http|https)://[a-zA-Z0-9\-._~:/?#[\]@!$&'()*+,;=]+$`

	match, err := regexp.MatchString(urlPattern, url)
	if err != nil {
		// Handle error if the regular expression compilation fails
		fmt.Println("Error:", err)
		return false
	}

	return match
}

// COMPANY

func NewCompany(name, candidatePortal string) (*Company, error) {
	// Candidate portal must be a valid URL
	if !isValidURL(candidatePortal) {
		return nil, fmt.Errorf("'%s' is not a valid URL", candidatePortal)
	}

	return &Company{
		Name:            name,
		CandidatePortal: candidatePortal,
	}, nil
}

func (c Company) String() string {
	return c.Name + ", " + c.CandidatePortal
}
