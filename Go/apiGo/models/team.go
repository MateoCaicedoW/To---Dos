package models

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

var (
	National = "national"
	Club     = "club"

	types = []string{National, Club}
)

type Team struct {
	ID      uuid.UUID `gorm:"primary_key; not null; unique"`
	Name    string
	Type    string
	Country string
}

func (t *Team) Validate() (response TeamResponse) {
	response.Data = ListTeams{}
	response.Status = http.StatusBadRequest
	nameTeam := strings.Replace(strings.ToLower(t.Name), " ", "", -1)
	typeTeam := strings.Replace(strings.ToLower(t.Type), " ", "", -1)
	countryTeam := strings.Replace(strings.ToLower(t.Country), " ", "", -1)
	if nameTeam == "" {
		response.Message = "Name cant not be empty."
		return
	}
	if t.numbersAndCaracters(nameTeam, "Name").Message != "" {
		response.Message = t.numbersAndCaracters(nameTeam, "Name").Message
		return
	}

	if typeTeam == "" {
		response.Message = "Type cant not be empty."
		return
	}
	if typeTeam == Club && countryTeam == "" {
		response.Message = "Country cant not be empty if Type is Club."
		return
	}
	if t.validateType().Message != "" {
		response.Message = t.validateType().Message
		return
	}
	if typeTeam != Club && countryTeam != "" {
		response.Message = "Country must be empty if Type is Seleccion."
		return
	}

	if t.numbersAndCaracters(countryTeam, "Country").Message != "" {
		response.Message = t.numbersAndCaracters(countryTeam, "Country").Message
		return

	}

	response.Message = ""
	return
}

func (t *Team) validateType() (response TeamResponse) {

	response.Status = http.StatusBadRequest
	typeTeam := strings.Replace(strings.ToLower(t.Type), " ", "", -1)
	for _, v := range types {
		if typeTeam == v {
			response.Message = ""
			return
		}
	}

	if t.numbersAndCaracters(typeTeam, "Type").Message != "" {
		response.Message = t.numbersAndCaracters(typeTeam, "Type").Message
		return
	}

	response.Message = "Type must be National or Club."
	return
}

func (p *Team) numbersAndCaracters(param string, field string) (response TeamResponse) {
	numbers := regexp.MustCompile("^[0-9]+$")
	caracters := regexp.MustCompile("^[!-/:-@[-`{-~-$]+$")
	if param != "" {
		for _, item := range strings.Split(param, "") {
			if caracters.MatchString(item) {
				response.Message = field + " cant not contains caracters."
				return
			}
			if numbers.MatchString(item) {
				response.Message = field + " cant not be a number."
				return
			}
		}
	}

	response.Message = ""
	return
}

type TeamResponse struct {
	Status  int
	Data    ListTeams
	Message string
}
type ListTeams []Team
