package actions

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mateo/apiGo/models"
)

func (handler handler) ListTeams(w http.ResponseWriter, r *http.Request) {
	var teams []models.Team
	var response models.TeamResponse
	w.Header().Set("Content-Type", "application/json")
	if result := handler.db.Find(&teams); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = result.Error.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusSeeOther)
	response.Status = http.StatusSeeOther
	response.Data = teams
	json.NewEncoder(w).Encode(response)
}

func (handler handler) ShowTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.TeamResponse
	params := mux.Vars(r)
	ID := params["id"]
	team, err := findTeam(handler, ID, w, response)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusAccepted)
	response.Data = models.ListTeams{team}
	response.Status = http.StatusAccepted
	json.NewEncoder(w).Encode(response)
}

func (handler handler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.TeamResponse
	var team models.Team
	json.NewDecoder(r.Body).Decode(&team)
	team.ID = uuid.New()
	team.Name = strings.Replace(strings.ToLower(team.Name), " ", "", -1)
	team.Type = strings.Replace(strings.ToLower(team.Type), " ", "", -1)
	team.Country = strings.Replace(strings.ToLower(team.Country), " ", "", -1)
	var teams []models.Team
	handler.db.Find(&teams)
	for _, t := range teams {
		if t.Name == team.Name && t.Country == team.Country && t.Type == team.Type {
			w.WriteHeader(http.StatusBadRequest)
			response.Status = http.StatusBadRequest
			response.Message = "Team already exists"
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	err := team.Validate()
	if err.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return

	}
	if result := handler.db.Create(&team); result.Error != nil {
		w.WriteHeader(http.StatusBadGateway)
		response.Message = result.Error.Error()
		response.Status = http.StatusBadGateway
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response.Status = http.StatusOK
	response.Data = models.ListTeams{team}
	json.NewEncoder(w).Encode(response)

}

func (handler handler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ID := params["id"]

	var response models.TeamResponse
	team, err := findTeam(handler, ID, w, response)
	if err != nil {
		return
	}
	if result := handler.db.Delete(&team); result.Error != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(response)
		return
	}
	var teams []models.Team
	if result := handler.db.Find(&teams); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = result.Error.Error()
		response.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response.Status = http.StatusOK
	response.Data = teams
	json.NewEncoder(w).Encode(response)

}

func (handler handler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	ID := params["id"]
	var response models.TeamResponse
	var tempUpdate models.Team
	json.NewDecoder(r.Body).Decode(&tempUpdate)
	team, err := findTeam(handler, ID, w, response)
	if err != nil {
		return
	}
	err2 := tempUpdate.Validate()
	if err2.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return

	}
	team.Name = tempUpdate.Name
	team.Country = tempUpdate.Country
	team.Type = tempUpdate.Type
	if result := handler.db.Save(&team); result.Error != nil {
		w.WriteHeader(http.StatusBadGateway)
		response.Message = result.Error.Error()
		response.Status = http.StatusBadGateway
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response.Status = http.StatusCreated
	response.Data = models.ListTeams{team}
	json.NewEncoder(w).Encode(response)

}

func findTeam(handler handler, ID string, w http.ResponseWriter, response models.TeamResponse) (team models.Team, err error) {

	if result := handler.db.First(&team, &ID); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		response.Message = team.Name + " not found"
		response.Status = http.StatusNotFound
		json.NewEncoder(w).Encode(response)
		err = errors.New("team not found")
		return
	}
	return
}
