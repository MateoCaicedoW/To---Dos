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

func TestShow(t *testing.T) {
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

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	key := req.FormValue("id")
	var player models.Player
	expected := h.db.First(&player, &key)
	if expected.Error != nil {
		t.Errorf("handler returned unexpected body: got %v want %v", "d54126be-9663-45aa-b135-c27d500c2f26", key)
	}
}

func TestPlayers(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	req, err := http.NewRequest("GET", "/players", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.ListPlayers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestCreate(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	newplayer := models.Player{
		ID:        uuid.New(),
		FirstName: "John",
		LastName:  "Riqui",
		Level:     2,
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

func TestUpdate(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	player := models.Player{
		FirstName: "Messi",
		LastName:  "Ronaldo",
		Level:     2,
	}

	jsonStr, err := json.Marshal(player)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPut, "/players/?id=f6cddd44-0e37-49d6-b947-0dee1baef9ed", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.UpdatePlayer)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDelete(t *testing.T) {
	DB := db.Init()
	h := New(DB)
	req, err := http.NewRequest("PUT", "/players/", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Set("id", "5e3564aa-5d12-4e5e-af7b-ba65f20cdc9e")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.DeletePlayer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
