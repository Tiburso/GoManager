package company

import (
	"testing"
)

func TestNewCompany(t *testing.T) {
	_, err := NewCompany("Test Company", "https://www.testcompany.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = NewCompany("Test Company", "Invalid URL")
	if err == nil {
		t.Fatalf("Expected error, got %v", err)
	}
}
