package actions

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mateo/apiGo/db"
	"github.com/mateo/apiGo/models"
)

func TestShowPlayer(t *testing.T) {
	DB := db.Test()
	h := New(DB)
	id := uuid.New().String()
	newPlayer := createPlayer(id)
	router := mux.NewRouter()
	router.HandleFunc("/players/{id}", h.ShowPlayer).Methods("GET")

	server := &http.Server{
		Addr:    ":2000",
		Handler: router,
	}
	requestresponse := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/players/"+newPlayer.ID.String(), nil)
	server.Handler.ServeHTTP(requestresponse, req)
	if status := requestresponse.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}

func TestListPlayers(t *testing.T) {
	DB := db.Test()
	h := New(DB)

	router := mux.NewRouter()
	router.HandleFunc("/players", h.ListPlayers).Methods("GET")

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	requestresponse := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/players", nil)
	server.Handler.ServeHTTP(requestresponse, req)
	if status := requestresponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}

}

func TestCreatePlayer(t *testing.T) {
	DB := db.Test()
	h := New(DB)
	newplayer := models.Player{
		FirstName:         "Ronaldo",
		LastName:          "Nazario",
		Level:             98,
		Age:               32,
		Position:          "Defender",
		PhysicalCondition: "A+",
		Teams:             []models.Team{},
	}

	jsonStr, err := json.Marshal(newplayer)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "/players", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	requestresponse := httptest.NewRecorder()

	handler := http.HandlerFunc(h.CreatePlayer)

	handler.ServeHTTP(requestresponse, req)
	if status := requestresponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdatePlayer(t *testing.T) {
	DB := db.Test()
	h := New(DB)
	id := uuid.New().String()
	newPlayer := createPlayer(id)

	player := models.Player{
		FirstName:         "Ronaldo",
		LastName:          "Nazario",
		Level:             2,
		Age:               32,
		Position:          "Defender",
		PhysicalCondition: "A+",
		Teams:             []models.Team{},
	}

	jsonStr, err := json.Marshal(player)
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/players/{id}", h.UpdatePlayer).Methods("PUT")

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	req, err := http.NewRequest(http.MethodPut, "/players/"+newPlayer.ID.String(), bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	log.Println(req.URL)
	requestresponse := httptest.NewRecorder()
	server.Handler.ServeHTTP(requestresponse, req)

	if status := requestresponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestDeletePlayer(t *testing.T) {
	DB := db.Test()
	h := New(DB)

	id := uuid.New().String()
	newPlayer := createPlayer(id)
	router := mux.NewRouter()
	router.HandleFunc("/players/{id}", h.DeletePlayer).Methods(http.MethodDelete)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	req, err := http.NewRequest(http.MethodDelete, "/players/"+newPlayer.ID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	requestresponse := httptest.NewRecorder()
	server.Handler.ServeHTTP(requestresponse, req)

	if status := requestresponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func createPlayer(id string) (newPlayer models.Player) {
	DB := db.Test()
	h := New(DB)

	newPlayer = models.Player{
		ID:                uuid.MustParse(id),
		FirstName:         "Mateo",
		LastName:          "Caicedo",
		Level:             98,
		Age:               32,
		Position:          "Defender",
		PhysicalCondition: "A+",
		Teams:             []models.Team{},
	}

	jsonStr, err := json.Marshal(newPlayer)
	if err != nil {
		panic(err)
	}
	req1, _ := http.NewRequest("POST", "/players", bytes.NewBuffer(jsonStr))

	req1.Header.Set("Content-Type", "application/json")
	request := httptest.NewRecorder()

	handler := http.HandlerFunc(h.CreatePlayer)

	handler.ServeHTTP(request, req1)

	return
}
