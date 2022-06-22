package actions

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mateo/apiGo/db"
	"github.com/mateo/apiGo/models"
)

func TestShowTeam(t *testing.T) {
	DB := db.Test()
	h := New(DB)
	id := uuid.New().String()
	rand.Seed(time.Now().UnixNano())
	newTeam := create(id, randomString(10), "National", "")

	router := mux.NewRouter()
	router.HandleFunc("/teams/{id}", h.ShowTeam).Methods("GET")

	server := &http.Server{
		Addr:    ":2000",
		Handler: router,
	}
	requestresponse := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/teams/"+newTeam.ID.String(), nil)
	server.Handler.ServeHTTP(requestresponse, req)

	if status := requestresponse.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}

}

func TestListTeams(t *testing.T) {
	DB := db.Test()
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
	DB := db.Test()
	h := New(DB)
	id := uuid.New()
	rand.Seed(time.Now().UnixNano())
	newTeam := models.Team{
		ID:      id,
		Name:    randomString(12),
		Type:    "National",
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
	DB := db.Test()
	h := New(DB)
	id := uuid.New().String()
	rand.Seed(time.Now().UnixNano())
	team := create(id, randomString(9), "National", "")

	newTeam := models.Team{
		Name:    randomString(9),
		Type:    "National",
		Country: "",
	}

	jsonStr, err := json.Marshal(newTeam)
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/teams/{id}", h.UpdateTeam).Methods("PUT")

	server := &http.Server{
		Addr:    ":2000",
		Handler: router,
	}
	req, err := http.NewRequest(http.MethodPut, "/teams/"+team.ID.String(), bytes.NewBuffer(jsonStr))
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

func TestDeleteTeam(t *testing.T) {
	DB := db.Test()
	h := New(DB)
	id := uuid.New()
	rand.Seed(time.Now().UnixNano())
	team := create(id.String(), randomString(13), "National", "")
	router := mux.NewRouter()
	router.HandleFunc("/teams/{id}", h.DeleteTeam).Methods(http.MethodDelete)

	server := &http.Server{
		Addr:    ":2000",
		Handler: router,
	}
	req, err := http.NewRequest(http.MethodDelete, "/teams/"+team.ID.String(), nil)
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

func create(id string, name string, t string, country string) (newTeam models.Team) {
	DB := db.Test()
	h := New(DB)
	newTeam = models.Team{
		ID:      uuid.MustParse(id),
		Name:    name,
		Type:    t,
		Country: country,
	}

	jsonStr, err := json.Marshal(newTeam)
	if err != nil {
		panic(err)
	}
	req, _ := http.NewRequest("POST", "/teams", bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(h.CreateTeam)

	handler.ServeHTTP(rr, req)
	return
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
