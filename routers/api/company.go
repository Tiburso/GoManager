package api

import (
	"encoding/json"
	"net/http"

	company_model "github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/models/db"
)

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Initialize the company variable
	company := &company_model.Company{}

	// Decode JSON from the request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(company); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Validate the company creation
	company, err := company_model.NewCompany(company.Name, company.CandidatePortal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Check if the company already exists
	res := db.DB.Limit(1).Find(&company_model.Company{}, "name = ?", company.Name)

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

	// Create the company in the db
	res = db.DB.Create(&company)
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

	var companies []company_model.Company
	db.DB.Find(&companies)

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(companies)
}

func GetCompanyWithApplicationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get company from the query name
	name := r.URL.Query().Get("name")

	// Get the company from the db
	var company company_model.Company
	res := db.DB.Model(&company_model.Company{}).Preload("Applications").First(&company, "name = ?", name)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(company)
}
