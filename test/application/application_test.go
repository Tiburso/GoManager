package application

import (
	"testing"

	a "github.com/Tiburso/GoManager/pkg/application"
)

func TestNewApplication(t *testing.T) {
	company, err := a.NewCompany("Test Company", "https://www.testcompany.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = a.NewApplication("Test Application", a.FullTime, "2021-01-01", *company)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = a.NewApplication("Test Application", "Invalid Type", "2021-01-01", *company)
	if err == nil {
		t.Fatalf("Expected error, got %v", err)
	}

	_, err = a.NewApplication("Test Application", a.FullTime, "Invalid Date", *company)
	if err == nil {
		t.Fatalf("Expected error, got %v", err)
	}
}

func TestSetStatus(t *testing.T) {
	company, err := a.NewCompany("Test Company", "https://www.testcompany.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	app, err := a.NewApplication("Test Application", a.FullTime, "2021-01-01", *company)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = app.SetStatus(a.Applied)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = app.SetStatus(a.Rejected)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = app.SetStatus(a.Accepted)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = app.SetStatus("Invalid Status")
	if err == nil {
		t.Fatalf("Expected error, got %v", err)
	}
}

func TestSetType(t *testing.T) {
	company, err := a.NewCompany("Test Company", "https://www.testcompany.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	app, err := a.NewApplication("Test Application", a.FullTime, "2021-01-01", *company)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = app.SetType(a.FullTime)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = app.SetType(a.PartTime)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = app.SetType(a.Internship)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = app.SetType("Invalid Type")
	if err == nil {
		t.Fatalf("Expected error, got %v", err)
	}
}

func TestSetApplicationDate(t *testing.T) {
	company, err := a.NewCompany("Test Company", "https://www.testcompany.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	app, err := a.NewApplication("Test Application", a.FullTime, "2021-01-01", *company)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = app.SetApplicationDate("2021-01-01")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = app.SetApplicationDate("Invalid Date")
	if err == nil {
		t.Fatalf("Expected error, got %v", err)
	}
}

func TestString(t *testing.T) {
	company, err := a.NewCompany("Test Company", "https://www.testcompany.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	app, err := a.NewApplication("Test Application", a.FullTime, "2021-01-01", *company)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Test Application, Full Time, Applied, 2021-01-01, https://www.testcompany.com"
	if app.String() != expected {
		t.Fatalf("Expected %s, got %s", expected, app.String())
	}
}
