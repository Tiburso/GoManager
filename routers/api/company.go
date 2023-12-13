package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiburso/GoManager/common/structs"
	company_service "github.com/Tiburso/GoManager/services/company"
	"github.com/Tiburso/GoManager/services/convert"
)

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Initialize the company variable
	company := &structs.Company{}

	// Decode JSON from the request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(company); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err := company_service.CreateCompany(
		company.Name,
		company.CandidatePortal,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusCreated)
}

func GetCompaniesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get all companies from the db
	companies, err := company_service.GetCompanies()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	companiesApi := convert.ToCompanies(companies)

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(companiesApi)
}

func GetCompanyWithApplicationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get company from the query name
	name := r.URL.Query().Get("name")

	// Get the company from the db
	company, err := company_service.GetCompanyWithApplications(name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	companyWithApplication := convert.ToCompanyWithApplications(company)

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(companyWithApplication)
}

func EditCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode JSON from the request body
	var company_struct structs.Company
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&company_struct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// call the service
	err := company_service.UpdateCompany(
		company_struct.Name,
		company_struct.CandidatePortal,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
}

func DeleteCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get company from the query name
	name := r.URL.Query().Get("name")

	// Delete the company from the db
	err := company_service.DeleteCompany(name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusNoContent)
}
