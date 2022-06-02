package models

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

var (
	National = "seleccion"
	Club     = "club"

	types = []string{National, Club}
)

type Team struct {
	ID      uuid.UUID `gorm:"primary_key"`
	Name    string
	Type    string
	Country string
}

func (t *Team) Validate() (response TeamResponse) {
	numbers := regexp.MustCompile("^[0-9]+$")
	response.Data = ListTeams{}
	response.Status = http.StatusBadRequest
	nameTeam := strings.Trim(strings.ToLower(t.Name), " ")
	typeTeam := strings.Trim(strings.ToLower(t.Type), " ")
	countryTeam := strings.Trim(strings.ToLower(t.Country), " ")
	if nameTeam == "" {
		response.Message = "Name cant not be empty."
		return
	}
	if nameTeam != "" && numbers.MatchString(nameTeam) {
		response.Message = "Name cant not be a number."
		return
	}

	if typeTeam == "" {
		response.Message = "Type cant not be empty"
		return
	}
	if typeTeam != "" && numbers.MatchString(typeTeam) {
		return
	}
	if typeTeam == Club && countryTeam == "" {
		response.Message = "Country cant not be empty if Type is Club"
		return
	}
	if countryTeam != "" && numbers.MatchString(countryTeam) {
		response.Message = "Country cant not be a number."
		return
	}
	if t.validateType().Message != "" {
		response.Message = t.validateType().Message
		return
	}
	if typeTeam != Club && countryTeam != "" {
		response.Message = "Country must be empty if Type is National"
		return
	}

	response.Message = ""
	return
}

func (t *Team) validateType() (response TeamResponse) {
	response.Data = nil
	response.Status = http.StatusBadRequest
	typeTeam := strings.Trim(strings.ToLower(t.Type), " ")
	for _, v := range types {
		if typeTeam == v {
			response.Message = ""
			return
		}
	}
	response.Message = "Type must be Seleccion or Club"
	return
}

type TeamResponse struct {
	Status  int
	Data    ListTeams
	Message string
}
type ListTeams []Team
