package actions

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/mateo/apiGo/models"
)

func (handler handler) SignPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var playerTeam models.PlayerTeam
	json.NewDecoder(r.Body).Decode(&playerTeam)
	var response models.PlayerResponse
	var response2 models.TeamResponse

	//validate fields
	if err := playerTeam.Validate(); err.Message != "" {
		response.Message = err.Message
		response.Status = err.Status
		json.NewEncoder(w).Encode(response)
		return
	}

	// find player
	player, err := findPlayer(handler, playerTeam.PlayerID.String(), w, response)
	if err != nil {
		return
	}

	//find team
	teams, err2 := findTeam(handler, playerTeam.TeamID.String(), w, response2)
	if err2 != nil {
		return
	}
	playerTeam.TeamID = teams.ID

	// validate if player is already in team
	response = validateCreate(player, playerTeam, teams, w)
	if response.Message != "" {
		return
	}

	// save
	if result := handler.db.Create(&playerTeam); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = result.Error.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	player.Teams = append(player.Teams, teams)
	response.Status = http.StatusOK
	response.Message = "Player was signed to " + teams.Name + " successfully."
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

func (handler handler) TransferPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.PlayerResponse
	var response2 models.TeamResponse

	//get data
	var playerTeam models.PlayerTeam
	json.NewDecoder(r.Body).Decode(&playerTeam)

	if err := playerTeam.Validate(); err.Message != "" {
		response.Message = err.Message
		response.Status = err.Status
		json.NewEncoder(w).Encode(response)
		return
	}
	//find team
	newTeam, err := findTeam(handler, playerTeam.TeamID.String(), w, response2)
	if err != nil {
		return
	}

	//find player
	player, err := findPlayer(handler, playerTeam.PlayerID.String(), w, response)
	if err != nil {
		return
	}

	//validate if player is already in team

	response, idOldTeam := validateUpdate(player, playerTeam, newTeam, w, newTeam)
	if response.Message != "" {
		return
	}

	// find team to remove player
	oldTeam, err := findTeam(handler, idOldTeam.String(), w, response2)
	if err != nil {

		return
	}

	// find playerTeam to remove
	var playerTeamToRemove models.PlayerTeam
	if result := handler.db.Where("player_id = ? AND team_id = ?", player.ID, oldTeam.ID).First(&playerTeamToRemove); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = result.Error.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	//set new team
	playerTeam.ID = playerTeamToRemove.ID
	playerTeam.TeamID = newTeam.ID
	playerTeam.PlayerID = player.ID
	log.Println(player.Teams)

	// save
	handler.db.Model(&playerTeamToRemove).Updates(playerTeam)

	w.WriteHeader(http.StatusOK)
	response.Status = http.StatusOK
	//response.Data = models.ListPlayers{player}
	response.Message = "Player was transfered from " + oldTeam.Name + " to " + newTeam.Name + " successfully."
	json.NewEncoder(w).Encode(&response)
}

func validateCreate(player models.Player, playerTeam models.PlayerTeam, teams models.Team, w http.ResponseWriter) (response models.PlayerResponse) {
	if len(player.Teams) > 0 {
		for _, team := range player.Teams {
			if team.ID == playerTeam.TeamID {
				log.Println()
				response.Message = "Player is already in this team."
				response.Status = http.StatusBadRequest
				json.NewEncoder(w).Encode(response)
				return
			}
			if team.Type == models.Club && teams.Type == models.Club {
				response.Message = "Player can't be in two clubs."
				response.Status = http.StatusBadRequest
				json.NewEncoder(w).Encode(response)
				return
			}
			if team.Type == models.National && teams.Type == models.National {
				response.Message = "Player can't be in two national teams."
				response.Status = http.StatusBadRequest
				json.NewEncoder(w).Encode(response)
				return
			}
		}
	}
	response.Message = ""
	return
}

func validateUpdate(player models.Player, playerTeam models.PlayerTeam, teams models.Team, w http.ResponseWriter,
	newTeam models.Team) (response models.PlayerResponse, idOldTeam uuid.UUID) {

	if len(player.Teams) == 0 {
		response.Message = "Player can not be transfered because he is not in any team."
		response.Status = http.StatusBadRequest
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	countclub := 0
	countnational := 0
	for _, item := range player.Teams {
		if item.Type == newTeam.Type {
			if idOldTeam == newTeam.ID {
				response.Message = "Player is already in " + newTeam.Name + "."
				response.Status = http.StatusBadRequest
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}
			idOldTeam = item.ID
		}
		if item.Type == newTeam.Type {
			if idOldTeam == newTeam.ID {
				response.Message = "Player is already in " + newTeam.Name + "."
				response.Status = http.StatusBadRequest
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}
			idOldTeam = item.ID
		}
		if item.Type == models.Club {
			countclub++
		}
		if item.Type == models.National {
			countnational++
		}
	}
	if newTeam.Type == models.Club && countclub == 0 {
		response.Message = "Player can not be transfered because he is not in any club. Please sign him in a club."
		response.Status = http.StatusBadRequest
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	if newTeam.Type == models.National && countnational == 0 {
		response.Message = "Player can not be transfered because he is not in any National team. Please sign him in a National team."
		response.Status = http.StatusBadRequest
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return

	}
	response.Message = ""
	return
}
