package models

import (
	"net/http"

	"github.com/google/uuid"
)

type Player struct {
	ID                uuid.UUID `gorm:"primary_key"`
	FirstName         string
	LastName          string
	Level             int64
	Edad              int64
	Position          string
	PhysicalCondition string
}

func (p *Player) Validate() (result TemplatePlayers) {
	if p.FirstName == "" {
		result = TemplatePlayers{Status: http.StatusBadRequest, Data: nil, Message: "You inserted an Int or string empty on FirstName"}
	} else if p.LastName == "" {
		result = TemplatePlayers{Status: http.StatusBadRequest, Data: nil, Message: "You inserted an Int or string empty on LastName"}
	} else if p.Level == 0 {
		result = TemplatePlayers{Status: http.StatusBadRequest, Data: nil, Message: "You inserted a string or 0 on Level"}
	} else if p.Edad == 0 {
		result = TemplatePlayers{Status: http.StatusBadRequest, Data: nil, Message: "You inserted a string or 0 on Edad"}
	} else if p.Position == "" {
		result = TemplatePlayers{Status: http.StatusBadRequest, Data: nil, Message: "You inserted an Int or string empty on Position"}
	} else if p.PhysicalCondition == "" {
		result = TemplatePlayers{Status: http.StatusBadRequest, Data: nil, Message: "You inserted an Int or string empty on PhysicalCondition"}
	} else {
		result = TemplatePlayers{Status: http.StatusBadRequest, Data: ListPlayers{}, Message: ""}
	}
	return
}

type TemplatePlayers struct {
	Status  int
	Data    ListPlayers
	Message string
}

type ListPlayers []Player
