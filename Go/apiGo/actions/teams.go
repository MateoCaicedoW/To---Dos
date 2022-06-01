package actions

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mateo/apiGo/models"
)

func (h handler) ListTeams(w http.ResponseWriter, r *http.Request) {
	var teams []models.Team
	w.Header().Set("Content-Type", "application/json")
	if result := h.db.Find(&teams); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error)
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(models.TemplateTeams{Status: 200, Data: teams, Message: ""})
	}

}

func (h handler) ShowTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.FormValue("id")
	var team models.Team
	if result := h.db.First(&team, &key); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&models.TemplateTeams{Status: 404, Data: models.ListTeams{}, Message: "Team was not found"})
		return
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(models.TemplateTeams{Status: 200, Data: models.ListTeams{team}, Message: ""})
		return
	}

}

func (h handler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var team models.Team
	json.NewDecoder(r.Body).Decode(&team)
	team.ID = uuid.New()
	err := team.Validate()
	if err.Data != nil {
		if result := h.db.Create(&team); result.Error != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(&models.TemplateTeams{Status: 502, Data: models.ListTeams{}, Message: "Team was not created"})
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Team was created successfully"))
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

}

func (h handler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := r.FormValue("id")
	var team models.Team
	if result := h.db.First(&team, &key); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&models.TemplateTeams{Status: 404, Data: models.ListTeams{}, Message: "Team was not found"})
		return
	} else {
		if result := h.db.Delete(&team); result.Error != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(&models.TemplateTeams{Status: 502, Data: models.ListTeams{}, Message: "Team was not created"})
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Team was deleted successfully"))
			return
		}
	}
}

func (h handler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.FormValue("id")

	var tempUpdate models.Team
	json.NewDecoder(r.Body).Decode(&tempUpdate)
	var team models.Team
	if result := h.db.First(&team, &key); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&models.TemplateTeams{Status: 404, Data: models.ListTeams{}, Message: "Team was not found"})
		return
	} else {
		err := tempUpdate.Validate()
		if err.Data != nil {
			team.Name = tempUpdate.Name
			team.Country = tempUpdate.Country
			team.Type = tempUpdate.Type
			if result := h.db.Save(&team); result.Error != nil {
				w.WriteHeader(http.StatusBadGateway)
				json.NewEncoder(w).Encode(&models.TemplateTeams{Status: 502, Data: models.ListTeams{}, Message: "Team was not updated"})
				return
			} else {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("Team was updated successfully"))
				return
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
	}
}
