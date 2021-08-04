package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RunController - Run the controller
func RunController() {
	// init router
	r := mux.NewRouter()
	r.HandleFunc("/api/url/short", MakeShorterURL).Methods("POST")
	r.HandleFunc("/api/url/short", GetAllURLs).Methods("GET")
	r.HandleFunc("/{hash}", ActualEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
