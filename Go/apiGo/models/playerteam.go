package models

import (
	"github.com/google/uuid"
)

type PlayerTeam struct {
	ID       int `gorm:"primary_key; AUTO_INCREMENT; not null;"`
	TeamID   uuid.UUID
	PlayerID uuid.UUID
}
