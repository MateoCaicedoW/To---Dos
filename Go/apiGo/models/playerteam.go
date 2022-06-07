package models

import (
	"github.com/google/uuid"
)

type PlayerTeam struct {
	ID       int `gorm:"primary_key;  auto_increment"`
	TeamID   uuid.UUID
	PlayerID uuid.UUID
}
