package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mateo/apiGo/actions"
)

func TestShow(t *testing.T) {
	req, err := http.NewRequest("GET", "/players", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(actions.Show)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// Check the response body is what we expect.

	expected := `{"Status":200,"Data":[{"ID":"deae98b1-feab-47d0-a64b-5808bb12d612","FirstName":"Juan","LastName":"Hernandez","Level":12},{"ID":"9a42e66f-f020-471b-893e-2e04c2e24307","FirstName":"Marcos","LastName":"Llorente","Level":90},{"ID":"07ec6cb3-71fa-458d-bd0a-afcf1f7c3e5d","FirstName":"Luis","LastName":"Avila","Level":82},{"ID":"18207fe0-c9af-4b98-a1cb-494d874450f0","FirstName":"Alberto","LastName":"Mercado","Level":20},{"ID":"5e98f402-d087-4134-9a01-59af9af0f6d1","FirstName":"Rosa","LastName":"Mercado","Level":30},{"ID":"dba0f115-c8e6-401f-a14b-9e47a01367b3","FirstName":"Marta","LastName":"Rosario","Level":30},{"ID":"f86166fc-1481-44ee-8435-e1e78529d25d","FirstName":"Isaias","LastName":"Perez","Level":40},{"ID":"642da712-0eb2-4e8f-a0b9-5f6f597eaaaf","FirstName":"Samuel","LastName":"Benitez","Level":45},{"ID":"1b82526c-1e7b-46d6-a47b-06e8e15784e4","FirstName":"Gonzalo","LastName":"Higuain","Level":25},{"ID":"ae205ca3-84a9-485e-8b94-949c21aac14b","FirstName":"Alberto","LastName":"Rosado","Level":75},{"ID":"f3cf2904-0780-4e4e-bc44-0b9f8f1babb6","FirstName":"Ismael","LastName":"Perez","Level":59}],"Message":""}`

	if rr.Body.String() != expected {

		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestShowID(t *testing.T) {
	req, err := http.NewRequest("GET", "/players", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	q := req.URL.Query()

	q.Set("id", "deae98b1-feab-47d0-a64b-5808bb12d612")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(actions.ShowID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"Status":200,"Data":[{"ID":"deae98b1-feab-47d0-a64b-5808bb12d612","FirstName":"Juan","LastName":"Hernandez","Level":12}],"Message":""}`
	if expected != rr.Body.String() {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
