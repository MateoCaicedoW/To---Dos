package main

import (
	"io/ioutil"
	"log"
	"net/http"

	cors "github.com/gorilla/handlers"
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

	//serve static files
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	//routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("./public/index.html")
		if err != nil {
			log.Println(err)
		}
		w.Write(file)
	})

	//routes to teams
	router.HandleFunc("/api/players", handler.ListPlayers).Methods(http.MethodGet)
	router.HandleFunc("/api/players/{id}", handler.ShowPlayer).Methods(http.MethodGet)
	router.HandleFunc("/api/players", handler.CreatePlayer).Methods(http.MethodPost)
	router.HandleFunc("/api/players/{id}", handler.DeletePlayer).Methods(http.MethodDelete)
	router.HandleFunc("/api/players/{id}", handler.UpdatePlayer).Methods(http.MethodPut)

	//routes to players
	router.HandleFunc("/api/teams", handler.ListTeams).Methods(http.MethodGet)
	router.HandleFunc("/api/teams/{id}", handler.ShowTeam).Methods(http.MethodGet)
	router.HandleFunc("/api/teams", handler.CreateTeam).Methods(http.MethodPost)
	router.HandleFunc("/api/teams/{id}", handler.DeleteTeam).Methods(http.MethodDelete)
	router.HandleFunc("/api/teams/{id}", handler.UpdateTeam).Methods(http.MethodPut)

	//routes to players_teams
	router.HandleFunc("/api/movements/sign-player", handler.SignPlayer).Methods(http.MethodPost)
	router.HandleFunc("/api/movements/transfer-player", handler.TransferPlayer).Methods(http.MethodPut)
	router.HandleFunc("/api/movements/unsign-player", handler.UnsignPlayer).Methods(http.MethodDelete)

	//route to positions
	router.HandleFunc("/api/positions", handler.ListPositions).Methods(http.MethodGet)
	router.HandleFunc("/api/conditions", handler.ListConditions).Methods(http.MethodGet)
	router.HandleFunc(("/api/types"), handler.ListTypes).Methods(http.MethodGet)

	//cors

	headers := cors.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := cors.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := cors.AllowedOrigins([]string{"*"})

	configCors := cors.CORS(headers, methods, origins)

	//configure server
	server := &http.Server{
		Addr:    ":3000",
		Handler: configCors(router),
	}

	//start server
	log.Println("Listen ....")
	server.ListenAndServe()

}
