package controllers

import (
	"apex-trading-platform/src/models"
	"encoding/json"
	"net/http"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := models.GetMetrics()
	json.NewEncoder(w).Encode(metrics)
}
