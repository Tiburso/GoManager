package application

import (
	"fmt"
	"regexp"
	"time"
)

const (
	FullTime   = "Full Time"
	PartTime   = "Part Time"
	Internship = "Internship"
)

type Company struct {
	Name            string `gorm:"primaryKey"`
	CandidatePortal string
}

type Application struct {
	Name            string `gorm:"primaryKey"`
	CompanyName     string `gorm:"primaryKey"`
	Type            string
	ApplicationDate time.Time
	Company         Company `gorm:"foreignKey:CompanyName;references:Name"`
}

// type User struct {
// 	gorm.Model
// 	CreditCard CreditCard `gorm:"foreignKey:UserName"`
// 	// use UserName as foreign key
// }

// type CreditCard struct {
// 	gorm.Model
// 	Number   string `gorm:"primaryKey"`
// 	UserName string
// }

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

// APPLICATION

func NewApplication(name, applicationType, applicationDate string, company Company) (*Application, error) {
	// Application type must be one of the following: Full Time, Part Time, Internship
	if applicationType != FullTime && applicationType != PartTime && applicationType != Internship {
		return nil, fmt.Errorf("'%s' is not a valid application type", applicationType)
	}

	// Application date must be a valid date
	date, err := time.Parse("2006-01-02", applicationDate)
	if err != nil {
		return nil, fmt.Errorf("'%s' is not a valid date", applicationDate)
	}

	return &Application{
		Name:            name,
		Type:            applicationType,
		ApplicationDate: date,
		Company:         company,
	}, nil
}

func (a *Application) SetName(name string) error {
	a.Name = name

	return nil
}

func (a *Application) SetType(applicationType string) error {
	if applicationType != FullTime && applicationType != PartTime && applicationType != Internship {
		return fmt.Errorf("'%s' is not a valid application type", applicationType)
	}

	a.Type = applicationType

	return nil
}

func (a *Application) SetApplicationDate(applicationDate string) error {
	date, err := time.Parse("2006-01-02", applicationDate)
	if err != nil {
		return fmt.Errorf("'%s' is not a valid date", applicationDate)
	}

	a.ApplicationDate = date

	return nil
}

func (a Application) String() string {
	return a.Name + ", " + a.Type + ", " + a.ApplicationDate.String() + ", " /* + a.Company.String() */
}
