package application

import (
	"fmt"
	"time"
)

const (
	FullTime   = "Full Time"
	PartTime   = "Part Time"
	Internship = "Internship"
)

type Application struct {
	Name            string `gorm:"primaryKey"`
	CompanyName     string `gorm:"primaryKey"`
	Type            string
	ApplicationDate time.Time
	Company         Company `gorm:"foreignKey:CompanyName;references:Name"`
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
	company_string := ""

	//if company exists
	if a.Company.Name != "" {
		company_string = ", " + a.Company.CandidatePortal
	}

	// for the time i only want the format of yyyy-mm-dd
	return a.Name + ", " + a.Type + ", " + a.ApplicationDate.Format("2006-01-02") + company_string
}
