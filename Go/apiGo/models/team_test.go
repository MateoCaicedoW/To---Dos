package models

import (
	"testing"

	"github.com/google/uuid"
)

func TestValidateTeam(t *testing.T) {
	newPlayers := Team{
		ID:      uuid.New(),
		Name:    "Junior",
		Type:    "Seleccion",
		Country: "",
	}
	err := newPlayers.Validate()

	if len(err.Data) != 0 || err.Status != 400 {
		t.Error(err)
	}
}
