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

func TestShowPlayer(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	req, err := http.NewRequest("GET", "/players/", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	q := req.URL.Query()

	q.Set("id", "d54126be-9663-45aa-b135-c27d500c2f26")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.ShowPlayer)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
	}

}

func TestListPlayers(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	req, err := http.NewRequest("GET", "/players", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.ListPlayers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusAccepted)
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
		Edad:              32,
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
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(h.CreatePlayer)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
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
		Edad:              32,
		Position:          "Defender",
		PhysicalCondition: "A+",
	}

	jsonStr, err := json.Marshal(player)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPut, "/players/?id=873d44af-6b8a-449d-82c7-cf5485955be9", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.UpdatePlayer)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestDeletePlayer(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	req, err := http.NewRequest("PUT", "/players/", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Set("id", "873d44af-6b8a-449d-82c7-cf5485955be9")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.DeletePlayer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
