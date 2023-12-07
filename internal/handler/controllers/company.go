package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Tiburso/GoManager/internal/application"
	"github.com/Tiburso/GoManager/internal/database"
)

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var company *application.Company
	var err error
	// get the body of our POST request
	json.NewDecoder(r.Body).Decode(company)

	// Validate the company creation
	company, err = application.NewCompany(company.Name, company.CandidatePortal)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	res := database.DB.Create(&company)

	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res.Error.Error())
		return
	}
}
