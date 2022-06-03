package actions

import (
	"encoding/json"
	"log"
	"net/http"

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
		response.Message = result.Error.Error()
		response.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}
	var result []models.PlayerTeam
	//h.db.Raw("SELECT * FROM player_team").Scan(&result)
	h.db.Table("player_team").Find(&result)
	//log.Println(result)
	for i, player := range players {
		idTempPlayer := player.ID.String()
		idResultPlayer := result[i].PlayerID.String()
		if idTempPlayer == idResultPlayer {
			idTeam := result[i].TeamID.String()
			var team []models.Team
			h.db.First(&team, "id = ?", idTeam)
			players[i].Teams = team
		}
	}

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
	var result []models.PlayerTeam
	//h.db.Raw("SELECT * FROM player_team").Scan(&result)
	h.db.Table("player_team").Find(&result)

	for _, item := range result {
		if item.PlayerID.String() == idPlayer {
			var team []models.Team
			h.db.First(&team, "id = ?", item.TeamID)
			player.Teams = team
		}
	}

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

	if result := h.db.Delete(&player); result.Error != nil {
		w.WriteHeader(http.StatusBadGateway)
		response.Message = result.Error.Error()
		response.Status = http.StatusBadGateway
		json.NewEncoder(w).Encode(response)
		return
	}
	var players []models.Player
	if result := h.db.Find(&players); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = result.Error.Error()
		response.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}
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

	var teamTemp []models.Team

	for i := range newPlayer.Teams {

		nameTeam := newPlayer.Teams[i].Name

		if result := h.db.First(&teamTemp, "name = ?", nameTeam); result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.Message = result.Error.Error()
			response.Status = http.StatusBadRequest
			json.NewEncoder(w).Encode(response)
			return
		}
		log.Println(teamTemp)
	}

	log.Println(teamTemp)
	//newPlayer.Teams = teamTemp

	// if result := h.db.Create(&newPlayer); result.Error != nil {
	// 	w.WriteHeader(http.StatusBadGateway)
	// 	response.Message = result.Error.Error()
	// 	response.Status = http.StatusBadGateway
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	response.Status = http.StatusOK
	response.Data = models.ListPlayers{newPlayer}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func (h handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.PlayerResponse
	params := mux.Vars(r)
	idPlayer := params["id"]

	var tempUpdate models.Player
	json.NewDecoder(r.Body).Decode(&tempUpdate)

	player, err := findPlayer(h, idPlayer, w, response)
	if err != nil {
		return
	}
	err2 := tempUpdate.Validate()
	if err2.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	if result := h.db.Model(&player).Updates(tempUpdate); result.Error != nil {
		w.WriteHeader(http.StatusBadGateway)
		response.Message = result.Error.Error()
		response.Status = http.StatusBadGateway
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusCreated)
	response.Status = http.StatusCreated
	response.Data = models.ListPlayers{player}
	json.NewEncoder(w).Encode(response)

}

// func response(data models.ListPlayers, status int, message string) (response models.PlayerResponse) {
// 	response.Data = data
// 	response.Status = status
// 	response.Message = message
// 	return
// }

func findPlayer(h handler, idPlayer string, w http.ResponseWriter, response models.PlayerResponse) (player models.Player, err error) {

	if result := h.db.First(&player, &idPlayer); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		response.Message = result.Error.Error()
		response.Status = http.StatusNotFound
		json.NewEncoder(w).Encode(response)
		err = result.Error
		return
	}

	return
}
