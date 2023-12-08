package api

import (
	"encoding/json"
	"net/http"

	"github.com/Tiburso/GoManager/models/application"
	"github.com/Tiburso/GoManager/models/company"
	"github.com/Tiburso/GoManager/models/db"
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

	// Check if the application already exists
	res := db.DB.Limit(1).Find(&application.Application{}, "name = ? AND company_name = ?", name, company_name)

	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	if res.RowsAffected > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Application already exists")
		return
	}

	// fetch the company from the db
	var company company.Company
	res = db.DB.First(&company, "name = ?", company_name)

	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	// Validate the application creation
	app, err := application.NewApplication(name, application_type, application_date, company)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Create the application in the db
	res = db.DB.Create(&app)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(app)
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
	res := db.DB.Where("name = ? AND company_name = ?", name, company_name).Delete(&application.Application{})

	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
}

func UpdateApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Initialize the application variable
	update_map := make(map[string]any)

	// Decode JSON from the request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&update_map); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	name, name_ok := update_map["name"]
	company_name, company_name_ok := update_map["company_name"]

	if !name_ok || !company_name_ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Missing name or company_name")
		return
	}

	// Update the application in the db
	var app application.Application
	res := db.DB.First(&app, "name = ? AND company_name = ?", name, company_name)

	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	// Now we have the application, we can update it
	status, status_ok := update_map["status"]

	if status_ok {
		err := app.SetStatus(status.(string))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Invalid status")
			return
		}
	}

	application_date, application_date_ok := update_map["application_date"]

	if application_date_ok {
		err := app.SetApplicationDate(application_date.(string))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Invalid application date")
			return
		}
	}

	application_type, application_type_ok := update_map["application_type"]

	if application_type_ok {
		err := app.SetType(application_type.(string))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Invalid application type")
			return
		}
	}

	res = db.DB.Save(&app)

	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(app)
}

func GetApplicationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var applications []application.Application
	db.DB.Preload("Company").Find(&applications)

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(applications)
}
