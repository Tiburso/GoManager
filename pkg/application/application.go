package application

import "time"

type Company struct {
	Name            string
	CandidatePortal string
}

type Application struct {
	Name            string
	Type            string
	ApplicationDate time.Time
	Company         Company
}
