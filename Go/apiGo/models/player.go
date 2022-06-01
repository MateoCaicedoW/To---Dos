package models

import "github.com/google/uuid"

type Player struct {
	ID        uuid.UUID `gorm:"primary_key"`
	FirstName string
	LastName  string
	Level     int64
}

func (p *Player) Validate() (result Template) {
	if p.FirstName == "" {
		result = Template{Status: 400, Data: ListPlayers{}, Message: "You inserted an Int or string empty on FirstName"}
	} else if p.LastName == "" {
		result = Template{Status: 400, Data: ListPlayers{}, Message: "You inserted an Int or string empty on LastName"}
	} else if p.Level == 0 {
		result = Template{Status: 400, Data: ListPlayers{}, Message: "You inserted a string or 0 on Level"}
	}
	return
}

type Template struct {
	Status  int
	Data    ListPlayers
	Message string
}

type ListPlayers []Player
