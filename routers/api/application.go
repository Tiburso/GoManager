package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiburso/GoManager/common/structs"
	"github.com/Tiburso/GoManager/services/company"
	"github.com/Tiburso/GoManager/services/convert"
)

func CreateApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Initialize the application variable
	updateMap := make(map[string]string)

	// Decode JSON from the request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updateMap); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Check if the required fields are present
	name, name_ok := updateMap["name"]
	company_name, company_name_ok := updateMap["company_name"]
	application_date, application_date_ok := updateMap["application_date"]
	application_type, application_type_ok := updateMap["type"]

	if !name_ok || !company_name_ok || !application_date_ok || !application_type_ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Missing name, company_name, application_date or application_type")
		return
	}

	// call the service
	err := company.CreateApplication(name, application_type, application_date, company_name)

	// TODO: need to check now if the error is duplicate app or missing company
	// for now just return the error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusCreated)
}

func DeleteApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Initialize the application variable
	vars := r.URL.Query()

	name, name_ok := vars["name"]
	company_name, company_name_ok := vars["company_name"]

	if !name_ok || !company_name_ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Missing name or company_name")
		return
	}

	// Delete the application from the db
	err := company.DeleteApplication(name[0], company_name[0])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode JSON from the request body
	var application_struct structs.Application
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&application_struct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if application_struct.Name == "" || application_struct.CompanyName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Missing name or company_name")
		return
	}

	// call the service
	err := company.UpdateApplication(
		application_struct.Name,
		application_struct.Type,
		application_struct.ApplicationDate,
		application_struct.Status,
		application_struct.CompanyName,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
}

func GetApplicationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	applications, err := company.GetApplications()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(convert.ToApplicationCreations(applications))
}
