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
	//routes to teams
	r.HandleFunc("/players", h.ListPlayers).Methods(http.MethodGet)
	r.HandleFunc("/players/", h.ShowPlayer).Queries("id", "{id}").Methods(http.MethodGet)
	r.HandleFunc("/players", h.CreatePlayer).Methods(http.MethodPost)
	r.HandleFunc("/players/", h.DeletePlayer).Queries("id", "{id}").Methods(http.MethodDelete)
	r.HandleFunc("/players/", h.UpdatePlayer).Queries("id", "{id}").Methods(http.MethodPut)
	//routes to players
	r.HandleFunc("/teams", h.ListTeams).Methods(http.MethodGet)
	r.HandleFunc("/teams/", h.ShowTeam).Queries("id", "{id}").Methods(http.MethodGet)
	r.HandleFunc("/teams", h.CreateTeam).Methods(http.MethodPost)
	r.HandleFunc("/teams/", h.DeleteTeam).Queries("id", "{id}").Methods(http.MethodDelete)
	r.HandleFunc("/teams/", h.UpdateTeam).Queries("id", "{id}").Methods(http.MethodPut)

	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	log.Println("Listen ....")
	server.ListenAndServe()
}
