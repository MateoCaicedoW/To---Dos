package actions

import (
	"encoding/json"
	"log"
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
		json.NewEncoder(w).Encode(models.Template{Status: 200, Data: players, Message: ""})
	}

}

func (h handler) Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	key := r.FormValue("id")
	log.Print(key)
	//log.Println(params["id"])
	var player models.Player
	if result := h.db.First(&player, &key); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.Template{Status: 400, Data: models.ListPlayers{}, Message: "Player was not found"})
		return
	} else {
		json.NewEncoder(w).Encode(models.Template{Status: 200, Data: models.ListPlayers{player}, Message: ""})
		return
	}

}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	key := r.FormValue("id")
	var player models.Player
	if result := h.db.First(&player, &key); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.Template{Status: 400, Data: models.ListPlayers{}, Message: "Player was not found"})
		return
	} else {
		if result := h.db.Delete(&player); result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&models.Template{Status: 400, Data: models.ListPlayers{}, Message: "Player was not found"})
			return
		} else {
			w.Write([]byte("Player created successfully"))
			return
		}
	}
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//players := models.P
	var newPlayers models.Player
	json.NewDecoder(r.Body).Decode(&newPlayers)
	newPlayers.ID = uuid.New()
	err := newPlayers.Validate()
	if err.Data != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	} else {
		if result := h.db.Create(&newPlayers); result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(result.Error)
		} else {
			w.Write([]byte("Player created successfully"))
		}

	}

}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	key := r.FormValue("id")

	var tempUpdate models.Player
	json.NewDecoder(r.Body).Decode(&tempUpdate)

	var player models.Player
	if result := h.db.First(&player, &key); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.Template{Status: 400, Data: models.ListPlayers{}, Message: "Player was not found"})
		return
	} else {
		err := tempUpdate.Validate()
		if err.Data != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			if result := h.db.Model(&player).Updates(tempUpdate); result.Error != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(&models.Template{Status: 400, Data: models.ListPlayers{}, Message: "Player was not found"})
				return
			} else {
				w.Write([]byte("Player updated successfully"))
				return
			}
		}
	}

}
