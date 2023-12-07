package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Tiburso/GoManager/internal/application"
	"github.com/Tiburso/GoManager/internal/database"
)

func CreateApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Initialize the application variable
	app := &application.Application{}

	// Decode JSON from the request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(app); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Validate the application creation
	app, err := application.NewApplication(app.Name, app.Type, app.ApplicationDate.String(), app.Company)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Check if the application already exists
	res := database.DB.Limit(1).Find(&application.Application{}, "name = ? AND company_name = ?", app.Name, app.CompanyName)

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

	// Create the application in the database
	res = database.DB.Create(&app)
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
	app := &application.Application{}

	// Decode JSON from the request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(app); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Delete the application from the database
	res := database.DB.Where("name = ? AND company_name = ?", app.Name, app.CompanyName).Delete(&application.Application{})

	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(app)
}

func UpdateApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Initialize the application variable
	app := &application.Application{}

	// Decode JSON from the request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(app); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// Update the application in the database
	res := database.DB.Where("name = ? AND company_name = ?", app.Name, app.CompanyName).Save(&app)

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
	database.DB.Preload("Company").Find(&applications)

	// Send a JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(applications)
}
