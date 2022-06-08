package models

import (
	"net/http"

	"github.com/google/uuid"
)

type PlayerTeam struct {
	ID       int `gorm:"primary_key; AUTO_INCREMENT; not null;"`
	TeamID   uuid.UUID
	PlayerID uuid.UUID
}

func (p *PlayerTeam) Validate() (response PlayerTeamResponse) {
	var list []string
	response.Status = http.StatusBadRequest
	response.Data = list

	if p.TeamID == uuid.Nil {
		response.Message = "TeamID can not be empty."
		return
	}

	if p.PlayerID == uuid.Nil {
		response.Message = "PlayerID can not be empty."
		return
	}

	response.Message = ""
	return
}

type PlayerTeamResponse struct {
	Data    []string
	Message string
	Status  int
}
