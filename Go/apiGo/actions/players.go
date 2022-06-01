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
		w.WriteHeader(http.StatusSeeOther)
		json.NewEncoder(w).Encode(models.TemplatePlayers{Status: http.StatusAccepted, Data: players, Message: ""})
	}

}

func (h handler) ShowPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.FormValue("id")
	var player models.Player
	if result := h.db.First(&player, &key); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&models.TemplatePlayers{Status: http.StatusNotFound, Data: models.ListPlayers{}, Message: "Player was not found"})
		return
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(models.TemplatePlayers{Status: http.StatusAccepted, Data: models.ListPlayers{player}, Message: ""})
		return
	}

}

func (h handler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := r.FormValue("id")
	var player models.Player
	if result := h.db.First(&player, &key); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&models.TemplatePlayers{Status: http.StatusNotFound, Data: models.ListPlayers{}, Message: "Player was not found"})
		return
	} else {
		if result := h.db.Delete(&player); result.Error != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(&models.TemplatePlayers{Status: http.StatusBadGateway, Data: models.ListPlayers{}, Message: "Player was not deleted"})
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Player was deleted successfully"))
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
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(&models.TemplatePlayers{Status: http.StatusBadGateway, Data: models.ListPlayers{}, Message: "Player was not created"})
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Player was created successfully"))
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return

	}

}

func (h handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.FormValue("id")

	var tempUpdate models.Player
	json.NewDecoder(r.Body).Decode(&tempUpdate)

	var player models.Player

	if result := h.db.First(&player, &key); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&models.TemplatePlayers{Status: http.StatusNotFound, Data: models.ListPlayers{}, Message: "Player was not found"})
		return
	} else {
		err := tempUpdate.Validate()
		if err.Data != nil {
			if result := h.db.Model(&player).Updates(tempUpdate); result.Error != nil {
				w.WriteHeader(http.StatusBadGateway)
				json.NewEncoder(w).Encode(&models.TemplatePlayers{Status: http.StatusBadGateway, Data: models.ListPlayers{}, Message: "Player was not updated"})
				return
			} else {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("Player was updated successfully"))
				return
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return

		}
	}

}
