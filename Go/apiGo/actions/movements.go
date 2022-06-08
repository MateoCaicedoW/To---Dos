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
		w.WriteHeader(http.StatusBadRequest)
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

	if err := validateCreate(player, playerTeam, teams, w); err.Message != "" {
		response.Message = err.Message
		response.Status = err.Status
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
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
	response.Message = player.FirstName + " " + player.LastName + " was signed to " + teams.Name + " successfully."
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
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
		w.WriteHeader(http.StatusBadRequest)
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

	exception, idOldTeam := validateUpdate(player, playerTeam, w, newTeam)
	if exception.Message != "" {
		response.Message = exception.Message
		response.Status = exception.Status
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	// find team to remove player
	oldTeam, err := findTeam(handler, idOldTeam.String(), w, response2)
	if err != nil {

		return
	}

	// find playerTeam to remove

	playerTeamToRemove := findPlayerTeamToRemove(handler, player.ID.String(), oldTeam.ID.String(), w, &response)

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
	response.Message = player.FirstName + " " + player.LastName + " was transfered from " + oldTeam.Name + " to " + newTeam.Name + " successfully."
	json.NewEncoder(w).Encode(response)
}

func (handler handler) UnsignPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.PlayerResponse
	var response2 models.TeamResponse

	//get data
	var playerTeam models.PlayerTeam
	json.NewDecoder(r.Body).Decode(&playerTeam)

	if err := playerTeam.Validate(); err.Message != "" {
		response.Message = err.Message
		response.Status = err.Status
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	//find team
	team, err := findTeam(handler, playerTeam.TeamID.String(), w, response2)
	if err != nil {
		return
	}

	//find player
	player, err := findPlayer(handler, playerTeam.PlayerID.String(), w, response)
	if err != nil {
		return
	}

	var playerTeamToRemove models.PlayerTeam

	for _, item := range player.Teams {
		if item.Type == team.Type {
			if item.ID == team.ID {
				// find playerTeam to remove
				playerTeamToRemove = findPlayerTeamToRemove(handler, player.ID.String(), team.ID.String(), w, &response)

			}

		}
	}
	//validate if player is already in team
	if playerTeamToRemove.ID == 0 {
		response.Message = player.FirstName + " " + player.LastName + " is not in this team."
		response.Status = http.StatusBadRequest
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// save
	handler.db.Delete(&playerTeamToRemove)
	w.WriteHeader(http.StatusOK)
	response.Status = http.StatusOK
	response.Message = player.FirstName + " " + player.LastName + " was unsigned from " + team.Name + " successfully."
	json.NewEncoder(w).Encode(response)
}

func validateCreate(player models.Player, playerTeam models.PlayerTeam, teams models.Team, w http.ResponseWriter) (response models.PlayerResponse) {
	response.Status = http.StatusBadRequest
	if len(player.Teams) > 0 {
		for _, team := range player.Teams {
			if team.ID == playerTeam.TeamID {
				response.Message = player.FirstName + " " + player.LastName + " is already in this team."

				return
			}
			if team.Type == models.Club && teams.Type == models.Club {
				response.Message = player.FirstName + " " + player.LastName + " can't be in two clubs."

				return
			}
			if team.Type == models.National && teams.Type == models.National {
				response.Message = player.FirstName + " " + player.LastName + " can't be in two national teams."

				return
			}
		}
	}

	response.Message = ""
	return
}

func validateUpdate(player models.Player, playerTeam models.PlayerTeam, w http.ResponseWriter,
	newTeam models.Team) (response models.PlayerResponse, idOldTeam uuid.UUID) {
	response.Status = http.StatusBadRequest

	if len(player.Teams) == 0 {
		response.Message = player.FirstName + " " + player.LastName + " " + "can not be transfered because he is not in any team."

		return
	}
	countclub := 0
	countnational := 0
	for _, item := range player.Teams {
		if item.Type == newTeam.Type {

			if item.ID == newTeam.ID {
				response.Message = player.FirstName + " " + player.LastName + " " + "is already in " + newTeam.Name + "."

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
		response.Message = player.FirstName + " " + player.LastName + " " + "can not be transfered because he is not in any club. Please sign him in a club."

		return
	}
	if newTeam.Type == models.National && countnational == 0 {
		response.Message = player.FirstName + " " + player.LastName + " " + "can not be transfered because he is not in any National team. Please sign him in a National team."

		return

	}
	response.Message = ""
	return
}

func findPlayerTeamToRemove(handler handler, IDPlayer string, IDTeam string, w http.ResponseWriter, response *models.PlayerResponse) (playerTeamToRemove models.PlayerTeam) {
	response.Status = http.StatusBadRequest

	if result := handler.db.Where("player_id = ? AND team_id = ?", IDPlayer, IDTeam).First(&playerTeamToRemove); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Message = result.Error.Error()
		json.NewEncoder(w).Encode(response)

		return
	}

	response.Message = ""
	return
}
