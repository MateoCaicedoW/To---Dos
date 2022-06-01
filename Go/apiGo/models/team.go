package models

import (
	"github.com/google/uuid"
)

type Team struct {
	ID      uuid.UUID `gorm:"primary_key"`
	Name    string
	Type    string
	Country string
}

func (t *Team) Validate() (result TemplateTeams) {
	if t.Name == "" {
		result = TemplateTeams{Status: 400, Data: nil, Message: "You inserted an Int or string empty on Name"}
	} else if t.Type == "" {
		result = TemplateTeams{Status: 400, Data: nil, Message: "You inserted an Int or string empty on Type"}
	} else if t.Type == "Club" && t.Country == "" {
		result = TemplateTeams{Status: 400, Data: nil, Message: "You inserted an Int or string empty on Country"}
	} else {
		result = TemplateTeams{Status: 400, Data: ListTeams{}, Message: ""}
	}
	return
}

type TemplateTeams struct {
	Status  int
	Data    ListTeams
	Message string
}
type ListTeams []Team
