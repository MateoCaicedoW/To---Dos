package actions

import (
	"encoding/json"
	"net/http"

	"github.com/mateo/apiGo/models"
)

func (handler handler) ListConditions(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)

	conditions := models.PhysicalCondition
	json, _ := json.Marshal(conditions)
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}
