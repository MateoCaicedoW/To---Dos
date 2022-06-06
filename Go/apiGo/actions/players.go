package actions

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mateo/apiGo/models"
)

func (h handler) ListPlayers(w http.ResponseWriter, r *http.Request) {
	var players []models.Player
	var response models.PlayerResponse
	w.Header().Set("Content-Type", "application/json")
	if result := h.db.Find(&players); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = strings.ToTitle(result.Error.Error())
		response.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}
	h.db.Preload("Teams").Find(&players)
	response.Status = http.StatusSeeOther
	response.Data = players
	w.WriteHeader(http.StatusSeeOther)
	json.NewEncoder(w).Encode(response)

}

func (h handler) ShowPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.PlayerResponse

	params := mux.Vars(r)
	idPlayer := params["id"]
	player, err := findPlayer(h, idPlayer, w, response)
	if err != nil {
		return
	}
	h.db.Preload("Teams").Find(&player)
	response.Status = http.StatusAccepted
	response.Data = models.ListPlayers{player}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)

}

func (h handler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.PlayerResponse
	params := mux.Vars(r)
	idPlayer := params["id"]
	player, err := findPlayer(h, idPlayer, w, response)
	if err != nil {
		return
	}

	h.db.Model(&player).Association("Teams").Clear()
	if result := h.db.Delete(&player); result.Error != nil {
		w.WriteHeader(http.StatusBadGateway)
		response.Message = strings.ToTitle(result.Error.Error())
		response.Status = http.StatusBadGateway
		json.NewEncoder(w).Encode(response)
		return
	}

	var players []models.Player
	if result := h.db.Find(&players); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = strings.ToTitle(result.Error.Error())
		response.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}
	h.db.Preload("Teams").Find(&players)
	w.WriteHeader(http.StatusOK)
	response.Status = http.StatusOK
	response.Data = players
	json.NewEncoder(w).Encode(response)

}

func (h handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.PlayerResponse
	var newPlayer models.Player
	json.NewDecoder(r.Body).Decode(&newPlayer)
	newPlayer.ID = uuid.New()

	err := newPlayer.Validate()
	if err.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	teams, response := findTeamPlayer(h, w, newPlayer)

	if response.Status != http.StatusOK {

		return
	}
	newPlayer.Teams = teams

	if result := h.db.Create(&newPlayer); result.Error != nil {
		w.WriteHeader(http.StatusBadGateway)
		response.Message = result.Error.Error()
		response.Status = http.StatusBadGateway
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Status = http.StatusOK
	response.Data = models.ListPlayers{newPlayer}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idPlayer := params["id"]
	w.Header().Set("Content-Type", "application/json")
	var response models.PlayerResponse
	var newPlayer models.Player
	json.NewDecoder(r.Body).Decode(&newPlayer)

	player, err2 := findPlayer(h, idPlayer, w, response)
	if err2 != nil {
		return
	}

	err := newPlayer.Validate()
	if err.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	teams, response := findTeamPlayer(h, w, newPlayer)

	if response.Status != http.StatusOK {
		return
	}
	newPlayer.Teams = teams
	//player.Teams = teams
	if result := h.db.Model(&player).Updates(newPlayer); result.Error != nil {
		w.WriteHeader(http.StatusBadGateway)
		response.Message = result.Error.Error()
		response.Status = http.StatusBadGateway
		json.NewEncoder(w).Encode(response)
		return
	}

	h.db.Model(&player).Association("Teams").Replace(teams)
	w.WriteHeader(http.StatusCreated)
	response.Status = http.StatusCreated
	response.Data = models.ListPlayers{player}
	json.NewEncoder(w).Encode(response)

}

func (h handler) allTeams() (teams []models.Team) {
	h.db.Find(&teams)
	return
}

func findPlayer(h handler, idPlayer string, w http.ResponseWriter, response models.PlayerResponse) (player models.Player, err error) {

	if result := h.db.First(&player, &idPlayer); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		response.Message = strings.ToTitle(result.Error.Error())
		response.Status = http.StatusNotFound
		json.NewEncoder(w).Encode(response)
		err = result.Error
		return
	}

	return
}

func findTeamPlayer(h handler, w http.ResponseWriter, tempUpdate models.Player) (teams []models.Team, response models.PlayerResponse) {
	var count int
	for i := range tempUpdate.Teams {

		nameTeam := strings.Replace(strings.ToLower(tempUpdate.Teams[i].Name), " ", "", -1)
		var teamTemp models.Team

		if result := h.db.Raw("SELECT * FROM teams WHERE name = ?", nameTeam).Scan(&teamTemp); result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.Message = strings.ToTitle(result.Error.Error())
			response.Status = http.StatusBadRequest
			json.NewEncoder(w).Encode(response)
			return
		}

		if len(teamTemp.Name) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			response.Message = strings.ToTitle(nameTeam) + " not found"
			response.Status = http.StatusBadRequest
			json.NewEncoder(w).Encode(response)
			return
		}
		allteams := h.allTeams()

		for _, item := range allteams {
			if item.Name == nameTeam {
				if strings.Replace(strings.ToLower(item.Type), " ", "", -1) == models.Club {
					count++
				}
			}
		}
		if count == 2 {
			w.WriteHeader(http.StatusBadRequest)
			response.Message = "Teams must be a club and a national team"
			response.Status = http.StatusBadRequest
			json.NewEncoder(w).Encode(response)
			return
		}

		teams = append(teams, teamTemp)

	}
	response.Status = http.StatusOK
	return
}
