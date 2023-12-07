package handler

import "github.com/gorilla/mux"

func SetupRouter() (r *mux.Router) {
	r = mux.NewRouter()

	// r.HandleFunc("/api/v1/applications", application.CreateApplication).Methods("POST")

	return r
}
