package actions

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mateo/apiGo/models"
)

func (h handler) ListPlayers(w http.ResponseWriter, r *http.Request) {
	var players []models.Player
	w.Header().Set("Content-Type", "application/json")
	if result := h.db.Find(&players); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error)
	} else {
		json.NewEncoder(w).Encode(models.TemplatePlayers{Status: 200, Data: players, Message: ""})
	}

}

func (h handler) ShowPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.FormValue("id")
	var player models.Player
	if result := h.db.First(&player, &key); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.TemplatePlayers{Status: 400, Data: models.ListPlayers{}, Message: "Player was not found"})
		return
	} else {
		json.NewEncoder(w).Encode(models.TemplatePlayers{Status: 200, Data: models.ListPlayers{player}, Message: ""})
		return
	}

}

func (h handler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := r.FormValue("id")
	var player models.Player
	if result := h.db.First(&player, &key); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.TemplatePlayers{Status: 400, Data: models.ListPlayers{}, Message: "Player was not found"})
		return
	} else {
		if result := h.db.Delete(&player); result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&models.TemplatePlayers{Status: 400, Data: models.ListPlayers{}, Message: "Player was not found"})
			return
		} else {
			w.Write([]byte("Player created successfully"))
			return
		}
	}
}

func (h handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newPlayers models.Player
	json.NewDecoder(r.Body).Decode(&newPlayers)
	newPlayers.ID = uuid.New()
	err := newPlayers.Validate()
	if err.Data != nil {
		if result := h.db.Create(&newPlayers); result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(result.Error)
		} else {
			w.Write([]byte("Player created successfully"))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)

	}

}

func (h handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.FormValue("id")

	var tempUpdate models.Player
	json.NewDecoder(r.Body).Decode(&tempUpdate)

	var player models.Player
	if result := h.db.First(&player, &key); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.TemplatePlayers{Status: 400, Data: models.ListPlayers{}, Message: "Player was not found"})
		return
	} else {
		err := tempUpdate.Validate()
		if err.Data != nil {
			if result := h.db.Model(&player).Updates(tempUpdate); result.Error != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(&models.TemplatePlayers{Status: 400, Data: models.ListPlayers{}, Message: "Player was not found"})
				return
			} else {
				w.Write([]byte("Player updated successfully"))
				return
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)

		}
	}

}
