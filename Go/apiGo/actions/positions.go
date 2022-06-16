package actions

import (
	"encoding/json"
	"net/http"

	"github.com/mateo/apiGo/models"
)

func (handler handler) ListPositions(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)

	positions := models.Positions
	json, _ := json.Marshal(positions)
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}
