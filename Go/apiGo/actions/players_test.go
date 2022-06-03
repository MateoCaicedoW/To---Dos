package actions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mateo/apiGo/db"
	"github.com/mateo/apiGo/models"
)

func TestShowPlayer(t *testing.T) {
	DB := db.Init()
	h := New(DB)

	router := mux.NewRouter()
	router.HandleFunc("/players/{id}", h.ShowPlayer).Methods("GET")

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	server.ListenAndServe()
	requestresponse := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/players/98cb12a0-e168-44c3-af5e-ab38b45ca84b", nil)
	server.Handler.ServeHTTP(requestresponse, req)
	if status := requestresponse.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}

func TestListPlayers(t *testing.T) {
	DB := db.Init()
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
	if status := requestresponse.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}

}

func TestCreatePlayer(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	newplayer := models.Player{
		ID:                uuid.New(),
		FirstName:         "John",
		LastName:          "Riqui",
		Level:             2,
		Age:               32,
		Position:          "Defender",
		PhysicalCondition: "A+",
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
	DB := db.Init()
	h := New(DB)
	player := models.Player{
		FirstName:         "Samuel",
		LastName:          "Solano",
		Level:             2,
		Age:               32,
		Position:          "Defender",
		PhysicalCondition: "A+",
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
	req, err := http.NewRequest(http.MethodPut, "/players/12a76c4f-42a1-4f9f-a013-aa8c65e16993", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	requestresponse := httptest.NewRecorder()
	server.Handler.ServeHTTP(requestresponse, req)

	if status := requestresponse.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestDeletePlayer(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/players/{id}", h.DeletePlayer).Methods(http.MethodDelete)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	req, err := http.NewRequest(http.MethodDelete, "/players/e99b98c6-0127-4ced-887a-66a9d89ef89c", nil)
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
