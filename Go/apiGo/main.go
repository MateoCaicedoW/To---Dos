package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateo/apiGo/actions"
	"github.com/mateo/apiGo/db"
)

func main() {
	DB := db.Init()
	h := actions.New(DB)

	r := mux.NewRouter()
	r.HandleFunc("/players", h.ListPlayers).Methods(http.MethodGet)
	r.HandleFunc("/players/", h.Show).Queries("id", "{id}").Methods(http.MethodGet)
	r.HandleFunc("/players", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/players/", h.Delete).Queries("id", "{id}").Methods(http.MethodDelete)
	r.HandleFunc("/players/", h.Update).Queries("id", "{id}").Methods(http.MethodPut)

	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	log.Println("Listen ....")
	server.ListenAndServe()
}
