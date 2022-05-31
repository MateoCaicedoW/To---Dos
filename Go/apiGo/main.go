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
	r.HandleFunc("/players/", actions.ShowID).Queries("id", "{id}").Methods(http.MethodGet)
	r.HandleFunc("/players", actions.Create).Methods(http.MethodPost)
	r.HandleFunc("/players/", actions.Delete).Queries("id", "{id}").Methods(http.MethodDelete)
	r.HandleFunc("/players/", actions.Update).Queries("id", "{id}").Methods(http.MethodPut)

	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	log.Println("Listen ....")
	server.ListenAndServe()
}
