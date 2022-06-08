package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateo/apiGo/actions"
	"github.com/mateo/apiGo/db"
)

func main() {
	//intance of db
	DB := db.Init()
	handler := actions.New(DB)

	//create router
	router := mux.NewRouter()

	//routes to teams
	router.HandleFunc("/players", handler.ListPlayers).Methods(http.MethodGet)
	router.HandleFunc("/players/{id}", handler.ShowPlayer).Methods(http.MethodGet)
	router.HandleFunc("/players", handler.CreatePlayer).Methods(http.MethodPost)
	router.HandleFunc("/players/{id}", handler.DeletePlayer).Methods(http.MethodDelete)
	router.HandleFunc("/players/{id}", handler.UpdatePlayer).Methods(http.MethodPut)

	//routes to players
	router.HandleFunc("/teams", handler.ListTeams).Methods(http.MethodGet)
	router.HandleFunc("/teams/{id}", handler.ShowTeam).Methods(http.MethodGet)
	router.HandleFunc("/teams", handler.CreateTeam).Methods(http.MethodPost)
	router.HandleFunc("/teams/{id}", handler.DeleteTeam).Methods(http.MethodDelete)
	router.HandleFunc("/teams/{id}", handler.UpdateTeam).Methods(http.MethodPut)

	//routes to players_teams
	router.HandleFunc("/movements/sign-player", handler.SignPlayer).Methods(http.MethodPost)
	router.HandleFunc("/movements/transfer-player", handler.TransferPlayer).Methods(http.MethodPut)
	router.HandleFunc("/movements/unsign-player", handler.UnsignPlayer).Methods(http.MethodDelete)

	//configure server
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Println("Listen ....")

	//start server
	server.ListenAndServe()

}
