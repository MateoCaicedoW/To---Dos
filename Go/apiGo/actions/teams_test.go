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

func TestShowTeam(t *testing.T) {
	DB := db.Init()
	h := New(DB)

	router := mux.NewRouter()
	router.HandleFunc("/teams/{id}", h.ShowTeam).Methods("GET")

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	server.ListenAndServe()
	requestresponse := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/teams/b471f072-6679-43da-bfdb-2144b001a5ee", nil)
	server.Handler.ServeHTTP(requestresponse, req)
	if status := requestresponse.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}

func TestListTeams(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	req, err := http.NewRequest("GET", "/teams", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.ListTeams)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
	}

}

func TestCreateTeam(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	newTeam := models.Team{
		IDTeam:  uuid.New(),
		Name:    "Italia",
		Type:    "Seleccion",
		Country: "",
	}

	jsonStr, err := json.Marshal(newTeam)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "/teams", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(h.CreateTeam)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdateTeam(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	newTeam := models.Team{
		Name:    "Germany",
		Type:    "Seleccion",
		Country: "",
	}

	jsonStr, err := json.Marshal(newTeam)
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/teams/{id}", h.UpdateTeam).Methods("PUT")

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	req, err := http.NewRequest(http.MethodPut, "/teams/98d643d5-2de6-486f-becf-ad23fd89d845", bytes.NewBuffer(jsonStr))
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

func TestDeleteTeam(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/teams/{id}", h.DeleteTeam).Methods(http.MethodDelete)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	req, err := http.NewRequest(http.MethodDelete, "/teams/98d643d5-2de6-486f-becf-ad23fd89d845", nil)
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
