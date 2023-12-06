package application

import (
	"testing"

	a "github.com/Tiburso/GoManager/pkg/application"
)

func TestNewCompany(t *testing.T) {
	_, err := a.NewCompany("Test Company", "https://www.testcompany.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = a.NewCompany("Test Company", "Invalid URL")
	if err == nil {
		t.Fatalf("Expected error, got %v", err)
	}
}
