package actions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/mateo/apiGo/models"
)

func TestShow(t *testing.T) {
	req, err := http.NewRequest("GET", "/players/", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	q := req.URL.Query()

	q.Set("id", "deae98b1-feab-47d0-a64b-5808bb12d612")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Show)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := models.P[0].ID
	if expected != uuid.MustParse("deae98b1-feab-47d0-a64b-5808bb12d612") {
		t.Errorf("handler returned unexpected body: got %v want %v", "deae98b1-feab-47d0-a64b-5808bb12d612", expected)
	}
}
func TestPlayers(t *testing.T) {
	req, err := http.NewRequest("GET", "/players", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListPlayers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// Check the response body is what we expect.

	expected := len(models.P)
	if expected != 11 {

		t.Errorf("handler returned unexpected body: got %v want %v", 11, expected)
	}
}

func TestCreate(t *testing.T) {
	var jsonStr = []byte(`{"ID":"5037efd5-9706-4164-b59d-f3d87992be61","FirstName":"xyz","LastName":"pqr","Level":10}`)
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

	handler := http.HandlerFunc(Create)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := len(models.P)
	if expected != 12 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			12, expected)
	}

}
func TestUpdate(t *testing.T) {

	player := models.Player{
		FirstName: "Messi",
		LastName:  "Ronaldo",
	}

	jsonStr, err := json.Marshal(player)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPut, "/players/?id=deae98b1-feab-47d0-a64b-5808bb12d612", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Update)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := models.P[0].FirstName
	if expected != "Messi" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			"Messi", expected)
	}
}

func TestDele(t *testing.T) {
	req, err := http.NewRequest("PUT", "/players/", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Set("id", "deae98b1-feab-47d0-a64b-5808bb12d612")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Delete)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := len(models.P)
	if expected != 11 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			11, expected)
	}
}
