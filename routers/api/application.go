package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiburso/GoManager/common/structs"
	typeconversions "github.com/Tiburso/GoManager/common/type_conversions"
	"github.com/Tiburso/GoManager/services/company"
	"github.com/Tiburso/GoManager/services/convert"
	"github.com/gorilla/mux"
)

func CreateApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	company_id, err := typeconversions.ConverToID(mux.Vars(r)["company_id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Initialize the application variable
	application := &structs.Application{}

	// Decode JSON from the request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(application); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// call the service
	app, err := company.CreateApplication(
		application.Name,
		application.Type,
		application.ApplicationDate,
		company_id)

	// TODO: need to check now if the error is duplicate app or missing company
	// for now just return the error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(convert.ToApplication(app))
}

func DeleteApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := typeconversions.ConverToID(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Delete the application from the db
	err = company.DeleteApplication(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := typeconversions.ConverToID(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Decode JSON from the request body
	var application_struct structs.Application
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&application_struct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// call the service
	err = company.UpdateApplication(
		id,
		application_struct.Name,
		application_struct.Type,
		application_struct.ApplicationDate,
		application_struct.Status,
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
	json.NewEncoder(w).Encode(convert.ToApplications(applications))
}
