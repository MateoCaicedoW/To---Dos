package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateo/apiGo/actions"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/players", actions.Show).Methods(http.MethodGet)
	r.HandleFunc("/players/{id}", actions.ShowID).Methods(http.MethodGet)
	r.HandleFunc("/players", actions.Create).Methods(http.MethodPost)
	r.HandleFunc("/players/{id}", actions.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/players/{id}", actions.Update).Methods(http.MethodPut)

	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	log.Println("Listen ....")
	server.ListenAndServe()
}
