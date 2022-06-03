package models

import "github.com/google/uuid"

type PlayerTeam struct {
	PlayerID uuid.UUID
	TeamID   uuid.UUID
}
