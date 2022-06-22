package actions

import (
	"net/http"

	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

func setupCorsResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
