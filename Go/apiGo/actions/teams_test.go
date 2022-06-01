package actions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/mateo/apiGo/db"
	"github.com/mateo/apiGo/models"
)

func TestShowTeam(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	req, err := http.NewRequest("GET", "/teams/", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	q := req.URL.Query()

	q.Set("id", "b471f072-6679-43da-bfdb-2144b001a5ee")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.ShowTeam)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
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
	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
	}

}

func TestCreateTeam(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	newTeam := models.Team{
		ID:      uuid.New(),
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
	req, err := http.NewRequest(http.MethodPut, "/teams/?id=c11a838c-0f5a-4131-bf1a-05645d3781e4", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.UpdateTeam)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestDeleteTeam(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	req, err := http.NewRequest("DELETE", "/teams/", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Set("id", "c11a838c-0f5a-4131-bf1a-05645d3781e4")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.DeleteTeam)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
