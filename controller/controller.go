package controller

import "github.com/gorilla/mux"

// RunController - Run the controller
func RunController() {

	// init router
	r := mux.NewRouter()

	r.HandleFunc("/api/url/short", MakeShorter).Methods("POST")
	r.HandleFunc("/api/url/short", GetShorter).Methods("GET")
}
