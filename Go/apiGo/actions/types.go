package actions

import (
	"encoding/json"
	"net/http"

	"github.com/mateo/apiGo/models"
)

func (handler handler) ListTypes(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)

	types := models.Types
	json, _ := json.Marshal(types)
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}
