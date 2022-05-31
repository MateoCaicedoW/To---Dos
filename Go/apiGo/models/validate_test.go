package models

import (
	"testing"

	"github.com/google/uuid"
)

func TestValidate(t *testing.T) {
	newPlayers := Player{
		ID:        uuid.New(),
		FirstName: "John",
		LastName:  "333",
		Level:     5,
	}
	err := newPlayers.Validate()

	if err.Data != nil {
		t.Errorf("Error")
	}
}
