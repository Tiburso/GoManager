package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Tiburso/GoManager/internal/application"
	"github.com/Tiburso/GoManager/internal/database"
)

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Initialize the company variable
	company := &application.Company{}

	// Decode JSON from the request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(company); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Validate the company creation
	company, err := application.NewCompany(company.Name, company.CandidatePortal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Check if the company already exists
	res := database.DB.Limit(1).Find(&application.Company{}, "name = ?", company.Name)

	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	if res.RowsAffected > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Company already exists")
		return
	}

	// Create the company in the database
	res = database.DB.Create(&company)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(company)
}

func GetCompaniesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var companies []application.Company
	database.DB.Find(&companies)

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(companies)
}

func GetCompanyWithApplicationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get company from the query name
	name := r.URL.Query().Get("name")

	// Get the company from the database
	var company application.Company
	res := database.DB.Model(&application.Company{}).Preload("Applications").First(&company, "name = ?", name)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(company)
}
