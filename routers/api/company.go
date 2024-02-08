package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiburso/GoManager/common/structs"
	typeconversions "github.com/Tiburso/GoManager/common/type_conversions"
	company_service "github.com/Tiburso/GoManager/services/company"
	"github.com/Tiburso/GoManager/services/convert"
	"github.com/gorilla/mux"
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

	c, err := company_service.CreateCompany(
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
	json.NewEncoder(w).Encode(convert.ToCompany(c))
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

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(convert.ToCompanies(companies))
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get company id from the URL and convert it to an uint
	id, err := typeconversions.ConverToID(mux.Vars(r)["id"])

	// If the conversion failed, return a 400
	//TODO: Check how to do this in a more modular way
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Get the company from the db
	company, err := company_service.GetCompany(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(convert.ToCompanyWithApplications(company))
}

func EditCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := typeconversions.ConverToID(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Decode JSON from the request body
	var company_struct structs.Company
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&company_struct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// call the service
	err = company_service.UpdateCompany(
		id,
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

	id, err := typeconversions.ConverToID(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Delete the company from the db
	err = company_service.DeleteCompany(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusNoContent)
}
