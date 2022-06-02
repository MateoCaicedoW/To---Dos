package models

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var (
	Aplus  = "a+"
	Aminus = "a-"
	Bplus  = "b+"
	Bminus = "b-"
	Cplus  = "c+"
	Cminus = "c-"
	Dplus  = "d+"
	Dminus = "d-"

	physicalConditions  = []string{Aplus, Aminus, Bplus, Bminus, Cplus, Cminus, Dplus, Dminus}
	GoalKeeper          = "portero"
	Defender            = "defensa"
	CentralMidfielder   = "mediocentro"
	Forward             = "delantero"
	FullBack            = "lateralderecho"
	HalfBack            = "lateralizquierdo"
	DefensiveMidfielder = "mediodefensivo"
	AttackingMidfielder = "medioofensivo"
	CentreBack          = "centro"
	Winger              = "extremo"

	positions = []string{GoalKeeper, Defender, CentralMidfielder, Forward, FullBack, HalfBack, DefensiveMidfielder, AttackingMidfielder, CentreBack, Winger}
)

type Player struct {
	ID                uuid.UUID `gorm:"primary_key"`
	FirstName         string
	LastName          string
	Level             int64
	Age               int64
	Position          string
	PhysicalCondition string
}

func (p *Player) Validate() (response PlayerResponse) {

	numbers := regexp.MustCompile("^[0-9]+$")
	letters := regexp.MustCompile("^[a-zA-Z]+$")
	response.Data = nil
	response.Status = http.StatusBadRequest
	if strings.Trim(p.FirstName, " ") == "" {
		response.Message = "FirstName cant not be empty."
		return
	}
	if strings.Trim(p.FirstName, " ") != "" && numbers.MatchString(p.FirstName) {
		response.Message = "FirstName cant not be a number."
		return
	}
	if strings.Trim(p.LastName, " ") != "" && numbers.MatchString(p.LastName) {
		response.Message = "LastName cant not be a number."
		return
	}
	if strings.Trim(p.LastName, " ") == "" {
		response.Message = "LastName cant not be empty."
		return
	}
	if p.Level < 1 || p.Level > 99 && letters.MatchString(strconv.Itoa(int(p.Level))) {
		response.Message = "Level must be a number."
		return
	}
	if p.Level < 1 || p.Level > 99 {
		response.Message = "Level must be between 1 and 99."
		return
	}
	if p.Age < 1 && letters.MatchString(strconv.Itoa(int(p.Age))) {
		response.Message = "Age must be a number."
		return
	}
	if p.Age < 1 {
		response.Message = "Age must be greater than 0."
	}
	if p.validatePosition().Message != "" {
		response.Message = p.validatePosition().Message
		return
	}
	if p.validatePhysicalCondition().Message != "" {
		response.Message = p.validatePhysicalCondition().Message
		return
	}
	response.Message = ""
	return

}

func (p *Player) validatePhysicalCondition() (response PlayerResponse) {
	numbers := regexp.MustCompile("^[0-9]+$")
	response.Data = nil
	response.Status = http.StatusBadRequest
	physicalCondition := strings.Trim(strings.ToLower(p.PhysicalCondition), " ")
	if physicalCondition == "" {
		response.Message = "PhysicalCondition cant not be empty."
		return
	}
	if physicalCondition != "" && numbers.MatchString(physicalCondition) {
		response.Message = "PhysicalCondition cant not be a number."
		return
	}

	for _, item := range physicalConditions {
		if item == physicalCondition {
			response.Message = ""
			return
		}
	}
	response.Message = "Insert a valid PhysicalCondition."
	return
}
func (p *Player) validatePosition() (response PlayerResponse) {
	numbers := regexp.MustCompile("^[0-9]+$")
	response.Data = nil
	response.Status = http.StatusBadRequest
	position := strings.Trim(strings.ToLower(p.Position), " ")
	if position == "" {
		response.Message = "Position cant not be empty."
		return

	}
	if position != "" && numbers.MatchString(position) {
		response.Message = "Position cant not be a number."
		return
	}

	for _, pos := range positions {
		if pos == position {
			response.Message = ""
			log.Println(pos, position)
			return
		}
	}
	response.Message = "Insert a valid Position."
	return
}

type PlayerResponse struct {
	Status  int
	Data    ListPlayers
	Message string
}

type ListPlayers []Player
